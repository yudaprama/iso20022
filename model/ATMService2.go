package model

// Withdrawal service provided by the ATM inside the session.
type ATMService2 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType1Code `xml:"SvcTp"`
}

func (a *ATMService2) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService2) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService2) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService2) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType1Code)(&value)
}
