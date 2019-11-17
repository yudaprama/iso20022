package model

// Inquiry service provided by the ATM inside the session.
type ATMService22 struct {

	// Unique identification of the customer session in which the transfer is performed.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of fund transfer selected by the customer or the ATM.
	ServiceType *ATMServiceType9Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService22) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService22) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService22) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType9Code)(&value)
}

func (a *ATMService22) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}
