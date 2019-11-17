package model

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse6 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse5) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse5) AddTransaction() *CardPaymentTransactionAdviceResponse6 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse6)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse5) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
