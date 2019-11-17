package model

// Information related to financial authorisation.
type AcquirerFinancialInitiation1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment1 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *CardTransactionContext1 `xml:"Cntxt"`

	// Card transaction for which the financial authorisation has been requested.
	Transaction *CardTransaction5 `xml:"Tx"`
}

func (a *AcquirerFinancialInitiation1) AddEnvironment() *CardTransactionEnvironment1 {
	a.Environment = new(CardTransactionEnvironment1)
	return a.Environment
}

func (a *AcquirerFinancialInitiation1) AddContext() *CardTransactionContext1 {
	a.Context = new(CardTransactionContext1)
	return a.Context
}

func (a *AcquirerFinancialInitiation1) AddTransaction() *CardTransaction5 {
	a.Transaction = new(CardTransaction5)
	return a.Transaction
}
