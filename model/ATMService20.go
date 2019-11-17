package model

// Context in which the transaction is performed.
type ATMService20 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType5Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService20) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService20) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService20) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType5Code)(&value)
}

func (a *ATMService20) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}
