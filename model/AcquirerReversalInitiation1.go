package model

// Information related to the reversal.
type AcquirerReversalInitiation1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment3 `xml:"Envt"`

	// Reversal transaction.
	Transaction *CardTransaction7 `xml:"Tx"`
}

func (a *AcquirerReversalInitiation1) AddEnvironment() *CardTransactionEnvironment3 {
	a.Environment = new(CardTransactionEnvironment3)
	return a.Environment
}

func (a *AcquirerReversalInitiation1) AddTransaction() *CardTransaction7 {
	a.Transaction = new(CardTransaction7)
	return a.Transaction
}
