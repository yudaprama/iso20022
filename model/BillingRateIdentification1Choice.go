package model

// Specifies the billing rate identification code.
type BillingRateIdentification1Choice struct {

	// Specifies the billing rate identification code, as defined in an external code list.
	Code *ExternalBillingRateIdentification1Code `xml:"Cd"`

	// Specifies the billing rate identification code, as defined in a proprietary format.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (b *BillingRateIdentification1Choice) SetCode(value string) {
	b.Code = (*ExternalBillingRateIdentification1Code)(&value)
}

func (b *BillingRateIdentification1Choice) SetProprietary(value string) {
	b.Proprietary = (*Max35Text)(&value)
}
