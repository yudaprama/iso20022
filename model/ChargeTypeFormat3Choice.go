package model

// Choice between a standard code or a proprietary code for specifying the type of charge.
type ChargeTypeFormat3Choice struct {

	// Standard code to specify the type of charge.
	Code *ChargeType17Code `xml:"Cd"`

	// Proprietary code for specifying the type of charge.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *ChargeTypeFormat3Choice) SetCode(value string) {
	c.Code = (*ChargeType17Code)(&value)
}

func (c *ChargeTypeFormat3Choice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
