package enums

type TransactionOutcome int

const (
	TRANSACTION_OUTCOME_SUCCESS TransactionOutcome = iota
	TRANSACTION_OUTCOME_FAILURE
	TRANSACTION_OUTCOME_SKIPPED
)

type NativeSpaceOutcome int

const (
	NATIVE_SPACE_SUCCESS                         NativeSpaceOutcome = iota
	NATIVE_SPACE_EXCEPTION_WITH_NONCE_BUMPING                       // gas fee charged
	NATIVE_SPACE_EXCEPTION_WITHOUT_NONCE_BUMPING                    // no gas fee charged
)

type EvmSpaceOutcome int

const (
	EVM_SPACE_FAIL EvmSpaceOutcome = iota
	EVM_SPACE_SUCCESS
	EVM_SPACE_SKIPPED = 0xff
)
