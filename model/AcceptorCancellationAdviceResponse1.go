package model

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse2 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse1) AddEnvironment() *CardPaymentEnvironment3 {
	a.Environment = new(CardPaymentEnvironment3)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse1) AddTransaction() *CardPaymentTransactionAdviceResponse2 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse2)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse1) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
