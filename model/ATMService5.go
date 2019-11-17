package model

// Inquiry service provided by the ATM inside the session.
type ATMService5 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType3Code `xml:"SvcTp"`
}

func (a *ATMService5) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService5) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService5) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType3Code)(&value)
}
