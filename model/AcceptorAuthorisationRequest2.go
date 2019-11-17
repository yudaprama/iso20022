package model

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment9 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction11 `xml:"Tx"`
}

func (a *AcceptorAuthorisationRequest2) AddEnvironment() *CardPaymentEnvironment9 {
	a.Environment = new(CardPaymentEnvironment9)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest2) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorAuthorisationRequest2) AddTransaction() *CardPaymentTransaction11 {
	a.Transaction = new(CardPaymentTransaction11)
	return a.Transaction
}
