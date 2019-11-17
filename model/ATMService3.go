package model

// Withdrawal service provided by the ATM inside the session.
type ATMService3 struct {

	// Unique identification of the service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of transaction selected by the customer.
	ServiceType *ATMServiceType2Code `xml:"SvcTp"`
}

func (a *ATMService3) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService3) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService3) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService3) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType2Code)(&value)
}
