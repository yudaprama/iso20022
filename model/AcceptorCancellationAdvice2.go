package model

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment18 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction16 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice2) AddEnvironment() *CardPaymentEnvironment18 {
	a.Environment = new(CardPaymentEnvironment18)
	return a.Environment
}

func (a *AcceptorCancellationAdvice2) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCancellationAdvice2) AddTransaction() *CardPaymentTransaction16 {
	a.Transaction = new(CardPaymentTransaction16)
	return a.Transaction
}
