package model

// Service provided by the ATM inside the session.
type ATMService16 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType3Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService16) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService16) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService16) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService16) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType3Code)(&value)
}

func (a *ATMService16) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}
