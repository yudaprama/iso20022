package model

// Cancellation request from an acceptor.
type AcceptorCancellationRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment12 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction15 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest2) AddEnvironment() *CardPaymentEnvironment12 {
	a.Environment = new(CardPaymentEnvironment12)
	return a.Environment
}

func (a *AcceptorCancellationRequest2) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorCancellationRequest2) AddTransaction() *CardPaymentTransaction15 {
	a.Transaction = new(CardPaymentTransaction15)
	return a.Transaction
}
