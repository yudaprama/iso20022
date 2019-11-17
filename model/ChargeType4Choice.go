package model

// Choice of formats for the type of charge.
type ChargeType4Choice struct {

	// Type of charge expressed as a code.
	Code *ChargeType12Code `xml:"Cd"`

	// Type of charge expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ChargeType4Choice) SetCode(value string) {
	c.Code = (*ChargeType12Code)(&value)
}

func (c *ChargeType4Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
