package model

// Cancellation response from the acquirer.
type AcceptorCancellationResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction35 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction27 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse3) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorCancellationResponse3) AddTransaction() *CardPaymentTransaction35 {
	a.Transaction = new(CardPaymentTransaction35)
	return a.Transaction
}

func (a *AcceptorCancellationResponse3) AddTransactionResponse() *CardPaymentTransaction27 {
	a.TransactionResponse = new(CardPaymentTransaction27)
	return a.TransactionResponse
}
