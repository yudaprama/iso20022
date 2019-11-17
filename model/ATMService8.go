package model

// Context in which the transaction is performed.
type ATMService8 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType5Code `xml:"SvcTp"`
}

func (a *ATMService8) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService8) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService8) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType5Code)(&value)
}
