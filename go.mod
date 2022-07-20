module github.com/Conflux-Chain/go-conflux-sdk

go 1.14

require (
	bou.ke/monkey v1.0.2
	github.com/ethereum/go-ethereum v1.10.17
	github.com/mcuadros/go-defaults v1.2.0
	github.com/openweb3/go-rpc-provider v0.2.2
	github.com/openweb3/go-sdk-common v0.0.0-20220720074746-a7134e1d372c
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.3.1
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.7.0
	gopkg.in/urfave/cli.v1 v1.20.0
	gotest.tools v2.2.0+incompatible

)

// replace github.com/openweb3/go-sdk-common => ../go-sdk-common
// replace github.com/ethereum/go-ethereum => ../../ethereum/go-ethereum
// replace github.com/openweb3/go-rpc-provider v0.2.0 => ../go-rpc-provider
