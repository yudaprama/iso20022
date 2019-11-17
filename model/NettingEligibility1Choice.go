package model

// Choice of format for the netting eligibility information.
type NettingEligibility1Choice struct {

	// Specifies whether the settlement transaction is eligible for netting.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Netting eligibility expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (n *NettingEligibility1Choice) SetIndicator(value string) {
	n.Indicator = (*YesNoIndicator)(&value)
}

func (n *NettingEligibility1Choice) AddProprietary() *GenericIdentification20 {
	n.Proprietary = new(GenericIdentification20)
	return n.Proprietary
}
