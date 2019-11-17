package model

// Choice of format for the beneficial ownership.
type BeneficialOwnership5Choice struct {

	// Specifies whether there is change of beneficial ownership.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Beneficial ownership information expressed a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (b *BeneficialOwnership5Choice) SetIndicator(value string) {
	b.Indicator = (*YesNoIndicator)(&value)
}

func (b *BeneficialOwnership5Choice) AddProprietary() *GenericIdentification47 {
	b.Proprietary = new(GenericIdentification47)
	return b.Proprietary
}
