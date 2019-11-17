package model

// Reconciliation response from the acquirer.
type AcceptorReconciliationResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment38 `xml:"Envt"`

	// Response from the acquirer to the reconciliation transaction.
	TransactionResponse *ResponseType5 `xml:"TxRspn"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation4 `xml:"Tx"`
}

func (a *AcceptorReconciliationResponse4) AddEnvironment() *CardPaymentEnvironment38 {
	a.Environment = new(CardPaymentEnvironment38)
	return a.Environment
}

func (a *AcceptorReconciliationResponse4) AddTransactionResponse() *ResponseType5 {
	a.TransactionResponse = new(ResponseType5)
	return a.TransactionResponse
}

func (a *AcceptorReconciliationResponse4) AddTransaction() *TransactionReconciliation4 {
	a.Transaction = new(TransactionReconciliation4)
	return a.Transaction
}
