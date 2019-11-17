package model

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment45 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext14 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction51 `xml:"Tx"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorAuthorisationRequest5) AddEnvironment() *CardPaymentEnvironment45 {
	a.Environment = new(CardPaymentEnvironment45)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest5) AddContext() *CardPaymentContext14 {
	a.Context = new(CardPaymentContext14)
	return a.Context
}

func (a *AcceptorAuthorisationRequest5) AddTransaction() *CardPaymentTransaction51 {
	a.Transaction = new(CardPaymentTransaction51)
	return a.Transaction
}

func (a *AcceptorAuthorisationRequest5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
