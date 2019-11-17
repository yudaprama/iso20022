package model

// Choice of format for the netting eligibility information.
type NettingEligibility4Choice struct {

	// Specifies whether the settlement transaction is eligible for netting.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Netting eligibility expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (n *NettingEligibility4Choice) SetIndicator(value string) {
	n.Indicator = (*YesNoIndicator)(&value)
}

func (n *NettingEligibility4Choice) AddProprietary() *GenericIdentification30 {
	n.Proprietary = new(GenericIdentification30)
	return n.Proprietary
}
