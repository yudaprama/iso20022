package model

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment22 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext6 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction25 `xml:"Tx"`
}

func (a *AcceptorCompletionAdvice3) AddEnvironment() *CardPaymentEnvironment22 {
	a.Environment = new(CardPaymentEnvironment22)
	return a.Environment
}

func (a *AcceptorCompletionAdvice3) AddContext() *CardPaymentContext6 {
	a.Context = new(CardPaymentContext6)
	return a.Context
}

func (a *AcceptorCompletionAdvice3) AddTransaction() *CardPaymentTransaction25 {
	a.Transaction = new(CardPaymentTransaction25)
	return a.Transaction
}
