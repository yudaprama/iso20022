package model

// Choice of formats for the type of charge.
type ChargeType5Choice struct {

	// Fee (charge/commission) expressed as a code.
	Code *InvestmentFundFee1Code `xml:"Cd"`

	// Fee (charge/commission) expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ChargeType5Choice) SetCode(value string) {
	c.Code = (*InvestmentFundFee1Code)(&value)
}

func (c *ChargeType5Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
