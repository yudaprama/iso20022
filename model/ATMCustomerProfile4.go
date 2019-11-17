package model

// Profile of the customer selected by an ATM.
type ATMCustomerProfile4 struct {

	// Describes the main way customer information was collected to build up the customer menu and to provide the service.
	RetrievalMode *ATMCustomerProfile1Code `xml:"RtrvlMd"`

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`
}

func (a *ATMCustomerProfile4) SetRetrievalMode(value string) {
	a.RetrievalMode = (*ATMCustomerProfile1Code)(&value)
}

func (a *ATMCustomerProfile4) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile4) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}
