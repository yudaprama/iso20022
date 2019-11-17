package model

// Choice of format for the beneficial ownership.
type BeneficialOwnership4Choice struct {

	// Specifies whether there is change of beneficial ownership.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Beneficial ownership information expressed a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (b *BeneficialOwnership4Choice) SetIndicator(value string) {
	b.Indicator = (*YesNoIndicator)(&value)
}

func (b *BeneficialOwnership4Choice) AddProprietary() *GenericIdentification30 {
	b.Proprietary = new(GenericIdentification30)
	return b.Proprietary
}
