package model

// Information related to the response of a reversal.
type AcquirerReversalResponse1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment4 `xml:"Envt"`

	// Reversal card transaction.
	Transaction *CardTransaction8 `xml:"Tx"`
}

func (a *AcquirerReversalResponse1) AddEnvironment() *CardTransactionEnvironment4 {
	a.Environment = new(CardTransactionEnvironment4)
	return a.Environment
}

func (a *AcquirerReversalResponse1) AddTransaction() *CardTransaction8 {
	a.Transaction = new(CardTransaction8)
	return a.Transaction
}
