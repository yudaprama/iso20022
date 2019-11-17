package model

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment1 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction1 `xml:"Tx"`
}

func (a *AcceptorAuthorisationRequest1) AddEnvironment() *CardPaymentEnvironment1 {
	a.Environment = new(CardPaymentEnvironment1)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest1) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorAuthorisationRequest1) AddTransaction() *CardPaymentTransaction1 {
	a.Transaction = new(CardPaymentTransaction1)
	return a.Transaction
}
