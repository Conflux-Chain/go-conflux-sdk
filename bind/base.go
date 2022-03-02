package bind

import (
	"context"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rlp"
)

// SignerFn is a signer function callback when a contract requires a method to
// sign the transaction before submission.
type SignerFn func(common.Address, *types.UnsignedTransaction) (*types.UnsignedTransaction, error)

// CallOpts is the collection of options to fine tune a contract call request.
type CallOpts struct {
	Pending     bool            // Whether to operate on the pending state or the last known one
	From        *types.Address  // Optional the sender address, otherwise the first account is used
	EpochNumber *types.Epoch    // Optional the block number on which the call should be performed
	Context     context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

type TransactOpts = types.ContractMethodSendOption

// FilterOpts is the collection of options to fine tune filtering for events
// within a bound contract.
type FilterOpts struct {
	Start *types.Epoch // Start of the queried range
	End   *types.Epoch // End of the range (nil = latest)

	// Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// WatchOpts is the collection of options to fine tune subscribing for events
// within a bound contract.
type WatchOpts struct {
	Start *types.Epoch // Start of the queried range (nil = latest)
	// Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// BoundContract is the base wrapper object that reflects a contract on the
// Ethereum network. It contains a collection of methods that are used by the
// higher level contract bindings to operate.
type BoundContract struct {
	address    types.Address      // Deployment address of the contract on the Ethereum blockchain
	abi        abi.ABI            // Reflect based ABI to access the correct Ethereum methods
	caller     ContractCaller     // Read interface to interact with the blockchain
	transactor ContractTransactor // Write interface to interact with the blockchain
	filterer   ContractFilterer   // Event filtering to interact with the blockchain
}

// NewBoundContract creates a low level contract interface through which calls
// and transactions may be made through.
func NewBoundContract(address types.Address, abi abi.ABI, caller ContractCaller, transactor ContractTransactor, filterer ContractFilterer) *BoundContract {
	return &BoundContract{
		address:    address,
		abi:        abi,
		caller:     caller,
		transactor: transactor,
		filterer:   filterer,
	}
}

// DeployContract deploys a contract onto the Ethereum blockchain and binds the
// deployment address with a Go wrapper.
func DeployContract(opts *TransactOpts, abi abi.ABI, bytecode []byte, backend ContractBackend, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, *BoundContract, error) {
	// Otherwise try to deploy the contract
	c := NewBoundContract(types.Address{}, abi, backend, backend, backend)

	input, err := c.abi.Pack("", params...)
	if err != nil {
		return nil, nil, nil, err
	}
	tx, hash, err := c.transact(opts, nil, append(bytecode, input...))
	if err != nil {
		return nil, nil, nil, err
	}

	return tx, hash, c, nil
}

// CreateAddress creates an ethereum address given the bytes and the nonce
func CreateAddress(b cfxaddress.Address, nonce uint64, initBytecode []byte) cfxaddress.Address {
	data, _ := rlp.EncodeToBytes([]interface{}{b, nonce, initBytecode})
	_addr := common.BytesToAddress(crypto.Keccak256(data)[12:])
	return cfxaddress.MustNew(_addr.String())
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (c *BoundContract) Call(opts *CallOpts, results *[]interface{}, method string, params ...interface{}) error {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(CallOpts)
	}
	if results == nil {
		results = new([]interface{})
	}
	// Pack the input, call and unpack the results
	input, err := c.abi.Pack(method, params...)
	inputStr := hexutil.Bytes(input).String()
	if err != nil {
		return err
	}

	var (
		msg    = types.CallRequest{From: opts.From, To: &c.address, Data: &inputStr}
		code   []byte
		output []byte
	)

	output, err = c.caller.Call(msg, opts.EpochNumber)
	if err != nil {
		return err
	}
	if len(output) == 0 {
		// Make sure we have a contract to operate on, and bail out otherwise.
		if code, err = c.caller.GetCode(c.address, opts.EpochNumber); err != nil {
			return err
		} else if len(code) == 0 {
			return ErrNoCode
		}
	}

	if len(*results) == 0 {
		res, err := c.abi.Unpack(method, output)
		*results = res
		return err
	}
	res := *results
	return c.abi.UnpackIntoInterface(res[0], method, output)
}

func (c *BoundContract) GenRequest(opts *CallOpts, method string, params ...interface{}) types.CallRequest {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(CallOpts)
	}
	// Pack the input, call and unpack the results
	input, _ := c.abi.Pack(method, params...)
	inputStr := hexutil.Bytes(input).String()

	msg := types.CallRequest{From: opts.From, To: &c.address, Data: &inputStr}
	return msg
}

func (c *BoundContract) DecodeOutput(results *[]interface{}, output []byte, method string) error {
	if len(*results) == 0 {
		res, err := c.abi.Unpack(method, output)
		*results = res
		return err
	}
	res := *results
	return c.abi.UnpackIntoInterface(res[0], method, output)
}

// Transact invokes the (paid) contract method with params as input values.
func (c *BoundContract) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	// Otherwise pack up the parameters and invoke the contract
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, nil, err
	}
	// todo(rjl493456442) check the method is payable or not,
	// reject invalid transaction at the first place
	return c.transact(opts, &c.address, input)
}

// RawTransact initiates a transaction with the given raw calldata as the input.
// It's usually used to initiate transactions for invoking **Fallback** function.
func (c *BoundContract) RawTransact(opts *TransactOpts, calldata []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	// todo(rjl493456442) check the method is payable or not,
	// reject invalid transaction at the first place
	return c.transact(opts, &c.address, calldata)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (c *BoundContract) Transfer(opts *TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	// todo(rjl493456442) check the payable fallback or receive is defined
	// or not, reject invalid transaction at the first place
	return c.transact(opts, &c.address, nil)
}

func (c *BoundContract) transact(opts *TransactOpts, contract *types.Address, input []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	utxBase := opts
	if opts == nil {
		utxBase = &TransactOpts{}
	}
	utx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase(*utxBase),
		To:                      contract,
		Data:                    types.NewBytes(input),
	}

	c.transactor.ApplyUnsignedTransactionDefault(&utx)

	hash, err := c.transactor.SendTransaction(utx)
	if err != nil {
		return nil, nil, err
	}

	return &utx, &hash, err
}

func (c *BoundContract) GenUnsignedTransaction(opts *TransactOpts, method string, params ...interface{}) types.UnsignedTransaction {
	input, _ := c.abi.Pack(method, params...)

	utxBase := opts
	if opts == nil {
		utxBase = &TransactOpts{}
	}

	return types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase(*utxBase),
		To:                      &c.address,
		Data:                    types.NewBytes(input),
	}
}

// FilterLogs filters contract logs for past blocks, returning the necessary
// channels to construct a strongly typed bound iterator on top of them.
func (c *BoundContract) FilterLogs(opts *FilterOpts, name string, query ...[]interface{}) (chan types.Log, error) {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(FilterOpts)
	}
	// Append the event selector to the query parameters and construct the topic set
	query = append([][]interface{}{{c.abi.Events[name].ID}}, query...)

	topics, err := abi.MakeTopics(query...)
	if err != nil {
		return nil, err
	}

	// Start the background filtering
	logs := make(chan types.Log, 128)

	config := types.LogFilter{
		Address:   []types.Address{c.address},
		Topics:    converHashesListToCfx(topics),
		FromEpoch: opts.Start,
	}
	if opts.End != nil {
		config.ToEpoch = opts.End
	}
	/* TODO(karalabe): Replace the rest of the method below with this when supported
	sub, err := c.filterer.SubscribeFilterLogs(ensureContext(opts.Context), config, logs)
	*/
	buff, err := c.filterer.GetLogs(config)
	if err != nil {
		return nil, err
	}

	for _, log := range buff {
		logs <- log
	}

	return logs, nil
}

// WatchLogs filters subscribes to contract logs for future blocks, returning a
// subscription object that can be used to tear down the watcher.
func (c *BoundContract) WatchLogs(opts *WatchOpts, name string, query ...[]interface{}) (chan types.SubscriptionLog, event.Subscription, error) {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(WatchOpts)
	}

	// Append the event selector to the query parameters and construct the topic set
	query = append([][]interface{}{{c.abi.Events[name].ID}}, query...)

	topics, err := abi.MakeTopics(query...)
	if err != nil {
		return nil, nil, err
	}
	// Start the background filtering
	logs := make(chan types.SubscriptionLog, 128)

	config := types.LogFilter{
		Address: []types.Address{c.address},
		Topics:  converHashesListToCfx(topics),
	}
	if opts.Start != nil {
		config.FromEpoch = opts.Start
	}
	sub, err := c.filterer.SubscribeLogs(logs, config)
	if err != nil {
		return nil, nil, err
	}
	return logs, sub, nil
}

// UnpackLog unpacks a retrieved log into the provided output structure.
func (c *BoundContract) UnpackLog(out interface{}, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := c.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	topicsInCommon := convertCfxHashesToCommon(log.Topics)
	return abi.ParseTopics(out, indexed, topicsInCommon[1:])
}

// UnpackLogIntoMap unpacks a retrieved log into the provided map.
func (c *BoundContract) UnpackLogIntoMap(out map[string]interface{}, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := c.abi.UnpackIntoMap(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	topicsInCommon := convertCfxHashesToCommon(log.Topics)
	return abi.ParseTopicsIntoMap(out, indexed, topicsInCommon[1:])
}

func converHashesListToCfx(topicsList [][]common.Hash) [][]types.Hash {
	topicsListInCfxHash := [][]types.Hash{}
	for _, topics := range topicsList {
		var topicsInCfxHash []types.Hash
		for _, vv := range topics {
			topicsInCfxHash = append(topicsInCfxHash, types.Hash(vv.String()))
		}
		topicsListInCfxHash = append(topicsListInCfxHash, topicsInCfxHash)
	}
	return topicsListInCfxHash
}

func convertCfxHashesToCommon(hashs []types.Hash) []common.Hash {
	topicsInCommon := []common.Hash{}
	for a := 0; a < len(hashs); a++ {
		topicInCommon := common.HexToHash(hashs[a].String())
		topicsInCommon = append(topicsInCommon, topicInCommon)
	}
	return topicsInCommon
}
