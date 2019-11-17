package model

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment2 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction7 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice1) AddEnvironment() *CardPaymentEnvironment2 {
	a.Environment = new(CardPaymentEnvironment2)
	return a.Environment
}

func (a *AcceptorCancellationAdvice1) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCancellationAdvice1) AddTransaction() *CardPaymentTransaction7 {
	a.Transaction = new(CardPaymentTransaction7)
	return a.Transaction
}
