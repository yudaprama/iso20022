package model

// Choice of format for the netting eligibility information.
type NettingEligibility5Choice struct {

	// Specifies whether the settlement transaction is eligible for netting.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Netting eligibility expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (n *NettingEligibility5Choice) SetIndicator(value string) {
	n.Indicator = (*YesNoIndicator)(&value)
}

func (n *NettingEligibility5Choice) AddProprietary() *GenericIdentification47 {
	n.Proprietary = new(GenericIdentification47)
	return n.Proprietary
}
