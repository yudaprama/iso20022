package model

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment47 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext15 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction55 `xml:"Tx"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorCompletionAdvice5) AddEnvironment() *CardPaymentEnvironment47 {
	a.Environment = new(CardPaymentEnvironment47)
	return a.Environment
}

func (a *AcceptorCompletionAdvice5) AddContext() *CardPaymentContext15 {
	a.Context = new(CardPaymentContext15)
	return a.Context
}

func (a *AcceptorCompletionAdvice5) AddTransaction() *CardPaymentTransaction55 {
	a.Transaction = new(CardPaymentTransaction55)
	return a.Transaction
}

func (a *AcceptorCompletionAdvice5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
