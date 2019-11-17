package model

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment34 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext9 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction40 `xml:"Tx"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorCompletionAdvice4) AddEnvironment() *CardPaymentEnvironment34 {
	a.Environment = new(CardPaymentEnvironment34)
	return a.Environment
}

func (a *AcceptorCompletionAdvice4) AddContext() *CardPaymentContext9 {
	a.Context = new(CardPaymentContext9)
	return a.Context
}

func (a *AcceptorCompletionAdvice4) AddTransaction() *CardPaymentTransaction40 {
	a.Transaction = new(CardPaymentTransaction40)
	return a.Transaction
}

func (a *AcceptorCompletionAdvice4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
