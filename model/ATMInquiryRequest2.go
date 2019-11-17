package model

// Information related to the request of an inquiry from an ATM.
type ATMInquiryRequest2 struct {

	// Environment in which the inquiry is performed.
	Environment *ATMEnvironment14 `xml:"Envt"`

	// Context in which the inquiry is performed.
	Context *ATMContext14 `xml:"Cntxt"`

	// Inquiry information for the transaction.
	Transaction *ATMTransaction29 `xml:"Tx"`
}

func (a *ATMInquiryRequest2) AddEnvironment() *ATMEnvironment14 {
	a.Environment = new(ATMEnvironment14)
	return a.Environment
}

func (a *ATMInquiryRequest2) AddContext() *ATMContext14 {
	a.Context = new(ATMContext14)
	return a.Context
}

func (a *ATMInquiryRequest2) AddTransaction() *ATMTransaction29 {
	a.Transaction = new(ATMTransaction29)
	return a.Transaction
}
