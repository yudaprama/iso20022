package model

// Cancellation request from an acceptor.
type AcceptorCancellationRequest5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment48 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext16 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction56 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest5) AddEnvironment() *CardPaymentEnvironment48 {
	a.Environment = new(CardPaymentEnvironment48)
	return a.Environment
}

func (a *AcceptorCancellationRequest5) AddContext() *CardPaymentContext16 {
	a.Context = new(CardPaymentContext16)
	return a.Context
}

func (a *AcceptorCancellationRequest5) AddTransaction() *CardPaymentTransaction56 {
	a.Transaction = new(CardPaymentTransaction56)
	return a.Transaction
}
