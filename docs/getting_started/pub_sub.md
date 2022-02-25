Subscribe
-------

go-conflux-sdk functional-reactive nature makes it really simple to setup observers that notify subscribers of events taking place on the blockchain.

To receive all new blocks head as they are added to the blockchain:

```golang
channel := make(chan types.BlockHeader, 100)
subscription, err := client.SubscribeNewHeads(channel)
```

To receive all new epochs as they are added to the blockchain:

```golang
channel := make(chan types.WebsocketEpochResponse, 100)
subscription, err := client.SubscribeEpochs(channel)
```

To receive all new logs as they are emited on the blockchain:

```golang
channel := make(chan types.SubscriptionLog, 100)
subscription, err := client.SubscribeLogs(channel, types.LogFilter{...})
```

The new happened event will be sent to channel, and an error signal will be sent to `subscription.Err()` when a subscription error occurs. Normally handle the channels like:
```golang
errorchan := sub.Err()
for {
	select {
	case err = <-errorchan:
		// handle when error
	case xxx := <-channel:
		// handle received signal
	}
}
```

Subscriptions should always be cancelled when no longer required:

```golang
subscription.unsubscribe();
```

It should be noted that when subscribing logs, a `SubscribeLogs` object is received. It has two fields `Log` and `ChainRerog`, one of them must be nil and the other not. When Log is not nil, it means that a Log is received. When field `ChainReorg` is not nil, that means chainreorg occurs. That represents the log related to epoch greater than or equal to `ChainReog.RevertTo` will become invalid, and the Dapp needs to be dealt with at the business level.

Please find Publish-Subscribe API documentation from https://developer.confluxnetwork.org/conflux-doc/docs/pubsub

