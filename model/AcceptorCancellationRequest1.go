package model

// Cancellation request from an acceptor.
type AcceptorCancellationRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment4 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction5 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest1) AddEnvironment() *CardPaymentEnvironment4 {
	a.Environment = new(CardPaymentEnvironment4)
	return a.Environment
}

func (a *AcceptorCancellationRequest1) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorCancellationRequest1) AddTransaction() *CardPaymentTransaction5 {
	a.Transaction = new(CardPaymentTransaction5)
	return a.Transaction
}
