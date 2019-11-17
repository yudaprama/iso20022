package model

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment49 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext17 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction59 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice5) AddEnvironment() *CardPaymentEnvironment49 {
	a.Environment = new(CardPaymentEnvironment49)
	return a.Environment
}

func (a *AcceptorCancellationAdvice5) AddContext() *CardPaymentContext17 {
	a.Context = new(CardPaymentContext17)
	return a.Context
}

func (a *AcceptorCancellationAdvice5) AddTransaction() *CardPaymentTransaction59 {
	a.Transaction = new(CardPaymentTransaction59)
	return a.Transaction
}
