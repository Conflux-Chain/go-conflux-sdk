package light

import (
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/light/contract"
	"github.com/Conflux-Chain/go-conflux-sdk/types/enums"
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var RelayInterval = 3 * time.Second

type EvmRelayConfig struct {
	LightNode  common.Address // light node contract
	LedgerInfo common.Address // ledger info contract
	Verifier   common.Address // MPT verification contract
	Admin      common.Address // management admin address or relayer

	EpochFrom uint64 // epoch for initialization
	GcLimits  int64  // maximum number of blocks to remove at a time

	GasLimit uint64 // Fixed gas limit to send transaction if specified
}

type EvmRelayer struct {
	EvmRelayConfig
	coreClient    *sdk.Client
	relayerClient *web3go.Client
	lightNode     *contract.LightNode
	txOpts        *bind.TransactOpts
	skippedRound  uint64
}

func NewEvmRelayer(coreClient *sdk.Client, relayerClient *web3go.Client, config EvmRelayConfig) *EvmRelayer {
	backend, signer := relayerClient.ToClientForContract()
	lightNode, err := contract.NewLightNode(config.LightNode, backend)
	if err != nil {
		panic(err.Error())
	}

	return &EvmRelayer{
		EvmRelayConfig: config,
		coreClient:     coreClient,
		relayerClient:  relayerClient,
		lightNode:      lightNode,
		txOpts: &bind.TransactOpts{
			From:     config.Admin,
			Signer:   signer,
			GasLimit: config.GasLimit,
		},
	}
}

func (r *EvmRelayer) Relay() {
	var initialized bool

	for {
		relayed, err := r.relay(initialized)
		if err != nil {
			logrus.WithError(err).Warn("Failed to relay")
		} else if relayed {
			if !initialized {
				initialized = true
			}

			continue
		}

		time.Sleep(RelayInterval)
	}
}

func (r *EvmRelayer) relay(initialized bool) (relayed bool, err error) {
	state, err := r.lightNode.State(nil)
	if err != nil {
		return false, errors.WithMessage(err, "Failed to get light node state")
	}

	// initialize light node
	if !initialized {
		if err = r.initLightNode(&state); err != nil {
			return false, errors.WithMessage(err, "Failed to initialize light node")
		}

		return true, nil
	}

	// relay pos block
	if relayed, err = r.relayPosBlock(&state); err != nil {
		return false, errors.WithMessage(err, "Failed to relay pos block")
	}

	if relayed {
		return true, nil
	}

	// garbage collect pow blocks
	if relayed, err = r.removePowBlocks(&state); err != nil {
		return false, errors.WithMessage(err, "Failed to remove pow blocks")
	}

	return relayed, nil
}

func (r *EvmRelayer) initLightNode(state *contract.ILightNodeState) error {
	if state.Epoch.Uint64() > 0 {
		logrus.Debug("Light node already initialized")
		return nil
	}

	if r.EpochFrom == 0 {
		logrus.Fatal("epoch not configured to initialize light node")
	}

	logrus.WithField("epoch", r.EpochFrom).Debug("Begin to initialize light node")

	// get committee from previous epoch
	lastEpochLedger, err := r.coreClient.Pos().GetLedgerInfoByEpoch(hexutil.Uint64(r.EpochFrom - 1))
	if err != nil {
		return errors.WithMessage(err, "Failed to get ledger of previous epoch")
	}

	committee, ok := contract.ConvertCommittee(lastEpochLedger)
	if !ok {
		logrus.Fatal("Committee not found")
	}

	// get ledger of first round
	ledger, err := r.coreClient.Pos().GetLedgerInfoByEpochAndRound(hexutil.Uint64(r.EpochFrom), 1)
	if err != nil {
		return errors.WithMessage(err, "Failed to get ledger")
	}

	if ledger == nil {
		logrus.Fatal("Ledger not found")
	}

	if ledger.LedgerInfo.CommitInfo.Pivot == nil {
		logrus.Fatal("Pivot in ledger is nil")
	}

	tx, err := r.lightNode.Initialize(r.txOpts,
		r.Admin, r.LedgerInfo, r.Verifier,
		committee, contract.ConvertLedger(ledger),
	)
	if err != nil {
		return errors.WithMessage(err, "Failed to send transaction")
	}

	if err = r.waitForSuccess(tx.Hash()); err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"epoch": r.EpochFrom,
		"round": 1,
		"pivot": uint64(ledger.LedgerInfo.CommitInfo.Pivot.Height),
	}).Info("Light node initialized")

	return nil
}

func (r *EvmRelayer) relayPosBlock(state *contract.ILightNodeState) (bool, error) {
	epoch := state.Epoch.Uint64()
	round := state.Round.Uint64() + 1
	if r.skippedRound > 0 {
		round = r.skippedRound + 1
	}

	committed, err := r.isCommitted(epoch, round)
	if err != nil {
		return false, errors.WithMessage(err, "Failed to check commitment status")
	}

	if !committed {
		logrus.WithField("epoch", epoch).WithField("round", round).Debug("No pos block to relay")
		return false, nil
	}

	logrus.WithField("epoch", epoch).WithField("round", round).Debug("Begin to relay pos block")

	ledger, err := r.coreClient.Pos().GetLedgerInfoByEpochAndRound(hexutil.Uint64(epoch), hexutil.Uint64(round))
	if err != nil {
		return false, errors.WithMessage(err, "Failed to get ledger by epoch and round")
	}

	// no ledger in round, just skip it
	if ledger == nil {
		logrus.WithField("epoch", epoch).WithField("round", round).Debug("No ledger info in this round")
		r.skippedRound = round
		return true, nil
	}

	pivot := ledger.LedgerInfo.CommitInfo.Pivot

	// both committee and pow pivot block unchanged
	if ledger.LedgerInfo.CommitInfo.NextEpochState == nil {
		if pivot == nil || uint64(pivot.Height) <= state.FinalizedBlockNumber.Uint64() {
			logrus.WithField("epoch", epoch).WithField("round", round).Debug("Pos block pivot not changed")
			r.skippedRound = round
			return true, nil
		}
	}

	// update committee or pivot block
	tx, err := r.lightNode.RelayPOS(r.txOpts, contract.ConvertLedger(ledger))
	if err != nil {
		return false, errors.WithMessage(err, "Failed to send transaction")
	}

	if err = r.waitForSuccess(tx.Hash()); err != nil {
		return false, err
	}

	logrus.WithFields(logrus.Fields{
		"epoch": epoch,
		"round": round,
		"pivot": uint64(pivot.Height),
	}).Info("Succeeded to relay pos block")

	r.skippedRound = 0

	return true, nil
}

func (r *EvmRelayer) isCommitted(epoch, round uint64) (bool, error) {
	status, err := r.coreClient.Pos().GetStatus()
	if err != nil {
		return false, errors.WithMessage(err, "Failed to get pos status")
	}

	block, err := r.coreClient.Pos().GetBlockByNumber(postypes.NewBlockNumber(uint64(status.LatestCommitted)))
	if err != nil {
		return false, errors.WithMessage(err, "Failed to get the latest committed block")
	}

	if block == nil {
		logrus.Fatal("Latest committed PoS block is nil")
	}

	logrus.WithFields(logrus.Fields{
		"epoch": uint64(block.Epoch),
		"round": uint64(block.Round),
	}).Debug("Latest committed block found")

	if epoch > uint64(block.Epoch) {
		return false, nil
	}

	if epoch < uint64(block.Epoch) {
		return true, nil
	}

	return round <= uint64(block.Round), nil
}

func (r *EvmRelayer) RelayPoWBlocks(headers [][]byte) error {
	tx, err := r.lightNode.RelayPOW(r.txOpts, headers)
	if err != nil {
		return errors.WithMessage(err, "Failed to send transaction")
	}

	return r.waitForSuccess(tx.Hash())
}

func (r *EvmRelayer) removePowBlocks(state *contract.ILightNodeState) (bool, error) {
	if state.Blocks.Cmp(state.MaxBlocks) <= 0 {
		return false, nil
	}

	tx, err := r.lightNode.RemoveBlockHeader(r.txOpts, big.NewInt(r.GcLimits))
	if err != nil {
		return false, errors.WithMessage(err, "Failed to send transaction")
	}

	if err = r.waitForSuccess(tx.Hash()); err != nil {
		return false, err
	}

	logrus.WithFields(logrus.Fields{
		"blocks": state.Blocks,
		"max":    state.MaxBlocks,
	}).Debug("Succeeded to remove PoW blocks")

	return true, nil
}

func (r *EvmRelayer) waitForSuccess(txHash common.Hash) error {
	time.Sleep(3 * time.Second)

	for {
		time.Sleep(time.Second)

		receipt, err := r.relayerClient.Eth.TransactionReceipt(txHash)
		if err != nil {
			logrus.WithError(err).Warn("Failed to wait for receipt")
		} else if receipt != nil {
			if uint8(*receipt.Status) == uint8(enums.EVM_SPACE_SUCCESS) {
				return nil
			}

			if receipt.TxExecErrorMsg != nil {
				return errors.Errorf("Transaction execution failed: %v", *receipt.TxExecErrorMsg)
			}

			return ErrTransactionExecutionFailed
		}
	}
}
