package model

// Reconciliation response from the acquirer.
type AcceptorReconciliationResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment7 `xml:"Envt"`

	// Response from the acquirer to the reconciliation transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation1 `xml:"Tx"`
}

func (a *AcceptorReconciliationResponse1) AddEnvironment() *CardPaymentEnvironment7 {
	a.Environment = new(CardPaymentEnvironment7)
	return a.Environment
}

func (a *AcceptorReconciliationResponse1) AddTransactionResponse() *ResponseType1 {
	a.TransactionResponse = new(ResponseType1)
	return a.TransactionResponse
}

func (a *AcceptorReconciliationResponse1) AddTransaction() *TransactionReconciliation1 {
	a.Transaction = new(TransactionReconciliation1)
	return a.Transaction
}
