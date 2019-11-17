package model

// Information related to the response of a financial authorisation.
type AcquirerFinancialResponse1 struct {

	// Environment of the transaction
	Environment *CardTransactionEnvironment2 `xml:"Envt"`

	// Context in which the card transaction is performed.
	Context *CardTransactionContext3 `xml:"Cntxt,omitempty"`

	// Card transaction for which the financial authorisation has been requested.
	Transaction *CardTransaction6 `xml:"Tx"`
}

func (a *AcquirerFinancialResponse1) AddEnvironment() *CardTransactionEnvironment2 {
	a.Environment = new(CardTransactionEnvironment2)
	return a.Environment
}

func (a *AcquirerFinancialResponse1) AddContext() *CardTransactionContext3 {
	a.Context = new(CardTransactionContext3)
	return a.Context
}

func (a *AcquirerFinancialResponse1) AddTransaction() *CardTransaction6 {
	a.Transaction = new(CardTransaction6)
	return a.Transaction
}
