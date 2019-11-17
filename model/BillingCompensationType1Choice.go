package model

// Defines the type of billing compensation.
type BillingCompensationType1Choice struct {

	// Defines the type of billing compensation, as published in an external billing compensation type code list.
	Code *ExternalBillingCompensationType1Code `xml:"Cd"`

	// Defines the type of billing compensation, as defined in a proprietary format.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (b *BillingCompensationType1Choice) SetCode(value string) {
	b.Code = (*ExternalBillingCompensationType1Code)(&value)
}

func (b *BillingCompensationType1Choice) SetProprietary(value string) {
	b.Proprietary = (*Max35Text)(&value)
}
