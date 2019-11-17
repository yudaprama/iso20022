package model

// Cancellation request from an acceptor.
type AcceptorCancellationRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment23 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction26 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest3) AddEnvironment() *CardPaymentEnvironment23 {
	a.Environment = new(CardPaymentEnvironment23)
	return a.Environment
}

func (a *AcceptorCancellationRequest3) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorCancellationRequest3) AddTransaction() *CardPaymentTransaction26 {
	a.Transaction = new(CardPaymentTransaction26)
	return a.Transaction
}
