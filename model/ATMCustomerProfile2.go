package model

// Profile of the customer selected by an ATM.
type ATMCustomerProfile2 struct {

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`
}

func (a *ATMCustomerProfile2) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile2) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}
