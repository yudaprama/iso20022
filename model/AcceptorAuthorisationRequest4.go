package model

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment32 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext8 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction36 `xml:"Tx"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorAuthorisationRequest4) AddEnvironment() *CardPaymentEnvironment32 {
	a.Environment = new(CardPaymentEnvironment32)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest4) AddContext() *CardPaymentContext8 {
	a.Context = new(CardPaymentContext8)
	return a.Context
}

func (a *AcceptorAuthorisationRequest4) AddTransaction() *CardPaymentTransaction36 {
	a.Transaction = new(CardPaymentTransaction36)
	return a.Transaction
}

func (a *AcceptorAuthorisationRequest4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
