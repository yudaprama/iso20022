package model

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment2 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction3 `xml:"Tx"`
}

func (a *AcceptorCompletionAdvice1) AddEnvironment() *CardPaymentEnvironment2 {
	a.Environment = new(CardPaymentEnvironment2)
	return a.Environment
}

func (a *AcceptorCompletionAdvice1) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCompletionAdvice1) AddTransaction() *CardPaymentTransaction3 {
	a.Transaction = new(CardPaymentTransaction3)
	return a.Transaction
}
