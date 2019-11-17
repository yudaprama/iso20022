package model

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment25 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation2 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest3) AddEnvironment() *CardPaymentEnvironment25 {
	a.Environment = new(CardPaymentEnvironment25)
	return a.Environment
}

func (a *AcceptorReconciliationRequest3) AddTransaction() *TransactionReconciliation2 {
	a.Transaction = new(TransactionReconciliation2)
	return a.Transaction
}
