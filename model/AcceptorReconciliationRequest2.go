package model

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment15 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation2 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest2) AddEnvironment() *CardPaymentEnvironment15 {
	a.Environment = new(CardPaymentEnvironment15)
	return a.Environment
}

func (a *AcceptorReconciliationRequest2) AddTransaction() *TransactionReconciliation2 {
	a.Transaction = new(TransactionReconciliation2)
	return a.Transaction
}
