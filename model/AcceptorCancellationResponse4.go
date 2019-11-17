package model

// Cancellation response from the acquirer.
type AcceptorCancellationResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction42 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction43 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse4) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorCancellationResponse4) AddTransaction() *CardPaymentTransaction42 {
	a.Transaction = new(CardPaymentTransaction42)
	return a.Transaction
}

func (a *AcceptorCancellationResponse4) AddTransactionResponse() *CardPaymentTransaction43 {
	a.TransactionResponse = new(CardPaymentTransaction43)
	return a.TransactionResponse
}
