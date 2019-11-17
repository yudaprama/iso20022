package model

// Choice of format for the CCP eligibility information.
type CentralCounterPartyEligibility4Choice struct {

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Central counterparty eligibility information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CentralCounterPartyEligibility4Choice) SetIndicator(value string) {
	c.Indicator = (*YesNoIndicator)(&value)
}

func (c *CentralCounterPartyEligibility4Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
