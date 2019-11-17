package model

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment50 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation4 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest5) AddEnvironment() *CardPaymentEnvironment50 {
	a.Environment = new(CardPaymentEnvironment50)
	return a.Environment
}

func (a *AcceptorReconciliationRequest5) AddTransaction() *TransactionReconciliation4 {
	a.Transaction = new(TransactionReconciliation4)
	return a.Transaction
}
