package model

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment36 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext11 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction44 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice4) AddEnvironment() *CardPaymentEnvironment36 {
	a.Environment = new(CardPaymentEnvironment36)
	return a.Environment
}

func (a *AcceptorCancellationAdvice4) AddContext() *CardPaymentContext11 {
	a.Context = new(CardPaymentContext11)
	return a.Context
}

func (a *AcceptorCancellationAdvice4) AddTransaction() *CardPaymentTransaction44 {
	a.Transaction = new(CardPaymentTransaction44)
	return a.Transaction
}
