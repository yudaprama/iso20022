package model

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse4 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse3) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse3) AddTransaction() *CardPaymentTransactionAdviceResponse4 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse4)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse3) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
