package model

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment20 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext5 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction22 `xml:"Tx"`
}

func (a *AcceptorAuthorisationRequest3) AddEnvironment() *CardPaymentEnvironment20 {
	a.Environment = new(CardPaymentEnvironment20)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest3) AddContext() *CardPaymentContext5 {
	a.Context = new(CardPaymentContext5)
	return a.Context
}

func (a *AcceptorAuthorisationRequest3) AddTransaction() *CardPaymentTransaction22 {
	a.Transaction = new(CardPaymentTransaction22)
	return a.Transaction
}
