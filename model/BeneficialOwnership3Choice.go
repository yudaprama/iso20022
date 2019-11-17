package model

// Choice of format for the beneficial ownership.
type BeneficialOwnership3Choice struct {

	// Specifies whether there is change of beneficial ownership.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Beneficial ownership information expressed a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (b *BeneficialOwnership3Choice) SetIndicator(value string) {
	b.Indicator = (*YesNoIndicator)(&value)
}

func (b *BeneficialOwnership3Choice) AddProprietary() *GenericIdentification38 {
	b.Proprietary = new(GenericIdentification38)
	return b.Proprietary
}
