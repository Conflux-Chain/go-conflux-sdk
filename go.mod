module github.com/Conflux-Chain/go-conflux-sdk

go 1.14

require (
	bou.ke/monkey v1.0.2
	github.com/BurntSushi/toml v0.3.1
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v0.0.0-20180603214616-504e848d77ea
	github.com/ethereum/go-ethereum v1.9.25
	github.com/golang/mock v1.4.3
	github.com/gorilla/websocket v1.4.1-0.20190629185528-ae1634f6a989
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/goconvey v1.6.4
	github.com/valyala/fasthttp v1.13.1
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
)

// replace github.com/ethereum/go-ethereum => ../../ethereum/go-ethereum
