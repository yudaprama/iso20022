package model

// Choice of format for the netting eligibility information.
type NettingEligibility3Choice struct {

	// Specifies whether the settlement transaction is eligible for netting.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Netting eligibility expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (n *NettingEligibility3Choice) SetIndicator(value string) {
	n.Indicator = (*YesNoIndicator)(&value)
}

func (n *NettingEligibility3Choice) AddProprietary() *GenericIdentification38 {
	n.Proprietary = new(GenericIdentification38)
	return n.Proprietary
}
