package model

// Choice of format for the CCP eligibility information.
type CentralCounterPartyEligibility5Choice struct {

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Central counterparty eligibility information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CentralCounterPartyEligibility5Choice) SetIndicator(value string) {
	c.Indicator = (*YesNoIndicator)(&value)
}

func (c *CentralCounterPartyEligibility5Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
