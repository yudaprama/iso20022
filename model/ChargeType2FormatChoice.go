package model

// Choice between a standard code or a proprietary code for specifying the type of charge.
type ChargeType2FormatChoice struct {

	// Standard code to specify the type of charge.
	Code *ChargeType14Code `xml:"Cd"`

	// Proprietary code for specifying the type of charge.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *ChargeType2FormatChoice) SetCode(value string) {
	c.Code = (*ChargeType14Code)(&value)
}

func (c *ChargeType2FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
