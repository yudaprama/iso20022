package model

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment10 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction13 `xml:"Tx"`
}

func (a *AcceptorCompletionAdvice2) AddEnvironment() *CardPaymentEnvironment10 {
	a.Environment = new(CardPaymentEnvironment10)
	return a.Environment
}

func (a *AcceptorCompletionAdvice2) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCompletionAdvice2) AddTransaction() *CardPaymentTransaction13 {
	a.Transaction = new(CardPaymentTransaction13)
	return a.Transaction
}
