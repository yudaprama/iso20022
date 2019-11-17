package model

// Information related to the authorisation initiation.
type AcquirerAuthorisationInitiation1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment1 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *CardTransactionContext1 `xml:"Cntxt"`

	// Card transaction for which the authorisation is requested.
	Transaction *CardTransaction15 `xml:"Tx"`
}

func (a *AcquirerAuthorisationInitiation1) AddEnvironment() *CardTransactionEnvironment1 {
	a.Environment = new(CardTransactionEnvironment1)
	return a.Environment
}

func (a *AcquirerAuthorisationInitiation1) AddContext() *CardTransactionContext1 {
	a.Context = new(CardTransactionContext1)
	return a.Context
}

func (a *AcquirerAuthorisationInitiation1) AddTransaction() *CardTransaction15 {
	a.Transaction = new(CardTransaction15)
	return a.Transaction
}
