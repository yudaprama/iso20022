package model

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse4 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse4) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse4) AddTransaction() *CardPaymentTransactionAdviceResponse4 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse4)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse4) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
