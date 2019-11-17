package model

// Cancellation response from the acquirer.
type AcceptorCancellationResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction6 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction10 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse1) AddEnvironment() *CardPaymentEnvironment3 {
	a.Environment = new(CardPaymentEnvironment3)
	return a.Environment
}

func (a *AcceptorCancellationResponse1) AddTransaction() *CardPaymentTransaction6 {
	a.Transaction = new(CardPaymentTransaction6)
	return a.Transaction
}

func (a *AcceptorCancellationResponse1) AddTransactionResponse() *CardPaymentTransaction10 {
	a.TransactionResponse = new(CardPaymentTransaction10)
	return a.TransactionResponse
}
