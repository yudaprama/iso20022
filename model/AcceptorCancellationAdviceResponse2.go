package model

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse3 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse2) AddEnvironment() *CardPaymentEnvironment11 {
	a.Environment = new(CardPaymentEnvironment11)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse2) AddTransaction() *CardPaymentTransactionAdviceResponse3 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse3)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse2) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
