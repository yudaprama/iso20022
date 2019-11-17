package model

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse3 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse2) AddEnvironment() *CardPaymentEnvironment11 {
	a.Environment = new(CardPaymentEnvironment11)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse2) AddTransaction() *CardPaymentTransactionAdviceResponse3 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse3)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse2) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
