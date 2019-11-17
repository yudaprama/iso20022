package model

// Information related to the response of an ATM inquiry from an ATM manager.
type ATMInquiryResponse1 struct {

	// Environment of the transaction.
	Environment *ATMEnvironment5 `xml:"Envt"`

	// Context in which the inquiry is performed.
	Context *ATMContext6 `xml:"Cntxt"`

	// Inquiry information for the transaction.
	Transaction *ATMTransaction7 `xml:"Tx"`
}

func (a *ATMInquiryResponse1) AddEnvironment() *ATMEnvironment5 {
	a.Environment = new(ATMEnvironment5)
	return a.Environment
}

func (a *ATMInquiryResponse1) AddContext() *ATMContext6 {
	a.Context = new(ATMContext6)
	return a.Context
}

func (a *ATMInquiryResponse1) AddTransaction() *ATMTransaction7 {
	a.Transaction = new(ATMTransaction7)
	return a.Transaction
}
