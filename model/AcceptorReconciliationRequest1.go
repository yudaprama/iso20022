package model

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment7 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation1 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest1) AddEnvironment() *CardPaymentEnvironment7 {
	a.Environment = new(CardPaymentEnvironment7)
	return a.Environment
}

func (a *AcceptorReconciliationRequest1) AddTransaction() *TransactionReconciliation1 {
	a.Transaction = new(TransactionReconciliation1)
	return a.Transaction
}
