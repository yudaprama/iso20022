package model

// Information related to the response to a reconciliation.
type AcquirerReconciliationResponse1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment5 `xml:"Envt"`

	// Reconciliation transaction.
	Transaction *CardTransaction10 `xml:"Tx"`
}

func (a *AcquirerReconciliationResponse1) AddEnvironment() *CardTransactionEnvironment5 {
	a.Environment = new(CardTransactionEnvironment5)
	return a.Environment
}

func (a *AcquirerReconciliationResponse1) AddTransaction() *CardTransaction10 {
	a.Transaction = new(CardTransaction10)
	return a.Transaction
}
