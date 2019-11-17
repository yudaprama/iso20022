package model

// Service provided by the ATM inside the session.
type ATMService4 struct {

	// Unique identification of the service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType5Code `xml:"SvcTp"`
}

func (a *ATMService4) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService4) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService4) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService4) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType5Code)(&value)
}
