package model

// Cancellation request from an acceptor.
type AcceptorCancellationRequest4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment35 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext10 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction41 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest4) AddEnvironment() *CardPaymentEnvironment35 {
	a.Environment = new(CardPaymentEnvironment35)
	return a.Environment
}

func (a *AcceptorCancellationRequest4) AddContext() *CardPaymentContext10 {
	a.Context = new(CardPaymentContext10)
	return a.Context
}

func (a *AcceptorCancellationRequest4) AddTransaction() *CardPaymentTransaction41 {
	a.Transaction = new(CardPaymentTransaction41)
	return a.Transaction
}
