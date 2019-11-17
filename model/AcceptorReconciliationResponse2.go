package model

// Reconciliation response from the acquirer.
type AcceptorReconciliationResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment19 `xml:"Envt"`

	// Response from the acquirer to the reconciliation transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation2 `xml:"Tx"`
}

func (a *AcceptorReconciliationResponse2) AddEnvironment() *CardPaymentEnvironment19 {
	a.Environment = new(CardPaymentEnvironment19)
	return a.Environment
}

func (a *AcceptorReconciliationResponse2) AddTransactionResponse() *ResponseType1 {
	a.TransactionResponse = new(ResponseType1)
	return a.TransactionResponse
}

func (a *AcceptorReconciliationResponse2) AddTransaction() *TransactionReconciliation2 {
	a.Transaction = new(TransactionReconciliation2)
	return a.Transaction
}
