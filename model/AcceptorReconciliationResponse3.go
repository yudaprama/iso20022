package model

// Reconciliation response from the acquirer.
type AcceptorReconciliationResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment38 `xml:"Envt"`

	// Response from the acquirer to the reconciliation transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation3 `xml:"Tx"`
}

func (a *AcceptorReconciliationResponse3) AddEnvironment() *CardPaymentEnvironment38 {
	a.Environment = new(CardPaymentEnvironment38)
	return a.Environment
}

func (a *AcceptorReconciliationResponse3) AddTransactionResponse() *ResponseType1 {
	a.TransactionResponse = new(ResponseType1)
	return a.TransactionResponse
}

func (a *AcceptorReconciliationResponse3) AddTransaction() *TransactionReconciliation3 {
	a.Transaction = new(TransactionReconciliation3)
	return a.Transaction
}
