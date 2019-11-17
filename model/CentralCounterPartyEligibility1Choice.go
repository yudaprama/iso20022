package model

// Choice of format for the CCP eligibility information.
type CentralCounterPartyEligibility1Choice struct {

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Central counterparty eligibility information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CentralCounterPartyEligibility1Choice) SetIndicator(value string) {
	c.Indicator = (*YesNoIndicator)(&value)
}

func (c *CentralCounterPartyEligibility1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
