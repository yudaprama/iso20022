package model

// Information related to the response of an authorisation.
type AcquirerAuthorisationResponse1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment2 `xml:"Envt"`

	// Context in which the card transaction is performed.
	Context *CardTransactionContext3 `xml:"Cntxt,omitempty"`

	// Card transaction for which the authorisation has been requested.
	Transaction *CardTransaction4 `xml:"Tx,omitempty"`
}

func (a *AcquirerAuthorisationResponse1) AddEnvironment() *CardTransactionEnvironment2 {
	a.Environment = new(CardTransactionEnvironment2)
	return a.Environment
}

func (a *AcquirerAuthorisationResponse1) AddContext() *CardTransactionContext3 {
	a.Context = new(CardTransactionContext3)
	return a.Context
}

func (a *AcquirerAuthorisationResponse1) AddTransaction() *CardTransaction4 {
	a.Transaction = new(CardTransaction4)
	return a.Transaction
}
