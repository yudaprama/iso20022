package model

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse1 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse1) AddEnvironment() *CardPaymentEnvironment3 {
	a.Environment = new(CardPaymentEnvironment3)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse1) AddTransaction() *CardPaymentTransactionAdviceResponse1 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse1)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse1) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
