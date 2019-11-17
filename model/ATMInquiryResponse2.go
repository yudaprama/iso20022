package model

// Information related to the response of an ATM inquiry from an ATM manager.
type ATMInquiryResponse2 struct {

	// Environment of the transaction.
	Environment *ATMEnvironment12 `xml:"Envt"`

	// Context in which the inquiry is performed.
	Context *ATMContext15 `xml:"Cntxt"`

	// Inquiry information for the transaction.
	Transaction *ATMTransaction21 `xml:"Tx"`
}

func (a *ATMInquiryResponse2) AddEnvironment() *ATMEnvironment12 {
	a.Environment = new(ATMEnvironment12)
	return a.Environment
}

func (a *ATMInquiryResponse2) AddContext() *ATMContext15 {
	a.Context = new(ATMContext15)
	return a.Context
}

func (a *ATMInquiryResponse2) AddTransaction() *ATMTransaction21 {
	a.Transaction = new(ATMTransaction21)
	return a.Transaction
}
