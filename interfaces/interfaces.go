package interfaces

import (
	"net/http"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

// ContractDeployResult for state change notification when deploying contract
type ContractDeployResult struct {
	//DoneChannel channel for notifying when contract deployed done
	DoneChannel      <-chan struct{}
	TransactionHash  *types.Hash
	Error            error
	DeployedContract Contractor
}

// HTTPRequester is interface for emitting a http requester
type HTTPRequester interface {
	Get(url string) (resp *http.Response, err error)
}
