package model

// Choice of format for the beneficial ownership.
type BeneficialOwnership1Choice struct {

	// Specifies whether there is change of beneficial ownership.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Beneficial ownership information expressed a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (b *BeneficialOwnership1Choice) SetIndicator(value string) {
	b.Indicator = (*YesNoIndicator)(&value)
}

func (b *BeneficialOwnership1Choice) AddProprietary() *GenericIdentification20 {
	b.Proprietary = new(GenericIdentification20)
	return b.Proprietary
}
