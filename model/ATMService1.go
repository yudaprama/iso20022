package model

// Withdrawal service provided by the ATM inside the session.
type ATMService1 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of transaction selected by the customer.
	ServiceType *ATMServiceType1Code `xml:"SvcTp"`
}

func (a *ATMService1) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService1) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService1) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType1Code)(&value)
}
