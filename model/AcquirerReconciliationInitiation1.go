package model

// Information related to the reconciliation.
type AcquirerReconciliationInitiation1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment5 `xml:"Envt"`

	// Reconciliation transaction.
	Transaction *CardTransaction9 `xml:"Tx"`
}

func (a *AcquirerReconciliationInitiation1) AddEnvironment() *CardTransactionEnvironment5 {
	a.Environment = new(CardTransactionEnvironment5)
	return a.Environment
}

func (a *AcquirerReconciliationInitiation1) AddTransaction() *CardTransaction9 {
	a.Transaction = new(CardTransaction9)
	return a.Transaction
}
