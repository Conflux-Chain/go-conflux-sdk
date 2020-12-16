package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type LocalizedBlockTrace struct {
	TransactionTraces []LocalizedTransactionTrace `json:"transactionTraces"`
}

type LocalizedTransactionTrace struct {
	Traces []LocalizedTrace `json:"traces"`
}

type LocalizedTrace struct {
	Type   string       `json:"type"`
	Action CallOrCreate `json:"action"`
}

type CallOrCreate struct {
	/// The sending account.
	From *Address `json:"from"`
	/// The destination account.
	To *Address `json:"to,omitempty"`
	/// The value transferred to the destination account.
	Value hexutil.Big `json:"value"`
	/// The gas available for executing the call.
	Gas hexutil.Big `json:"gas"`
	/// The input data provided to the call.
	Input hexutil.Bytes `json:"input,omitempty"`
	/// The type of the call.
	CallType string `json:"callType,omitempty"`
	/// The init code.
	Init hexutil.Bytes `json:"init,omitempty"`
}

// UnmarshalJSON unmarshals Input and Init type from []byte to hexutil.Bytes
func (c *CallOrCreate) UnmarshalJSON(data []byte) error {
	type Alias CallOrCreate
	aux := &struct {
		Input []byte `json:"input"`
		Init  []byte `json:"init"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.Input = aux.Input
	c.Init = aux.Init
	return nil
}
