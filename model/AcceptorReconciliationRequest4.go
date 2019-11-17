package model

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment37 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation3 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest4) AddEnvironment() *CardPaymentEnvironment37 {
	a.Environment = new(CardPaymentEnvironment37)
	return a.Environment
}

func (a *AcceptorReconciliationRequest4) AddTransaction() *TransactionReconciliation3 {
	a.Transaction = new(TransactionReconciliation3)
	return a.Transaction
}
