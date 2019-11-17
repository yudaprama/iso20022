package model

// Choice of format for the CCP eligibility information.
type CentralCounterPartyEligibility3Choice struct {

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Central counterparty eligibility information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (c *CentralCounterPartyEligibility3Choice) SetIndicator(value string) {
	c.Indicator = (*YesNoIndicator)(&value)
}

func (c *CentralCounterPartyEligibility3Choice) AddProprietary() *GenericIdentification38 {
	c.Proprietary = new(GenericIdentification38)
	return c.Proprietary
}
