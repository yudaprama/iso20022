package model

// Profile of the customer selected by an ATM.
type ATMCustomerProfile1 struct {

	// Describes the main way customer information was collected to build up the customer menu and the withdrawal request.
	RetrievalMode *ATMCustomerProfile1Code `xml:"RtrvlMd"`

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`
}

func (a *ATMCustomerProfile1) SetRetrievalMode(value string) {
	a.RetrievalMode = (*ATMCustomerProfile1Code)(&value)
}

func (a *ATMCustomerProfile1) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile1) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}
