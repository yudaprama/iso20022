package model

// Cancellation response from the acquirer.
type AcceptorCancellationResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction6 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction10 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse2) AddEnvironment() *CardPaymentEnvironment11 {
	a.Environment = new(CardPaymentEnvironment11)
	return a.Environment
}

func (a *AcceptorCancellationResponse2) AddTransaction() *CardPaymentTransaction6 {
	a.Transaction = new(CardPaymentTransaction6)
	return a.Transaction
}

func (a *AcceptorCancellationResponse2) AddTransactionResponse() *CardPaymentTransaction10 {
	a.TransactionResponse = new(CardPaymentTransaction10)
	return a.TransactionResponse
}
