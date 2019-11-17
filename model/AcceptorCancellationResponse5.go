package model

// Cancellation response from the acquirer.
type AcceptorCancellationResponse5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction57 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction58 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse5) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorCancellationResponse5) AddTransaction() *CardPaymentTransaction57 {
	a.Transaction = new(CardPaymentTransaction57)
	return a.Transaction
}

func (a *AcceptorCancellationResponse5) AddTransactionResponse() *CardPaymentTransaction58 {
	a.TransactionResponse = new(CardPaymentTransaction58)
	return a.TransactionResponse
}
