package model

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment24 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction28 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice3) AddEnvironment() *CardPaymentEnvironment24 {
	a.Environment = new(CardPaymentEnvironment24)
	return a.Environment
}

func (a *AcceptorCancellationAdvice3) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCancellationAdvice3) AddTransaction() *CardPaymentTransaction28 {
	a.Transaction = new(CardPaymentTransaction28)
	return a.Transaction
}
