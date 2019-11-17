package model

// Specifies the charge type.
type ChargeType2Choice struct {

	// Charge type, in a coded form.
	Code *ChargeType1Code `xml:"Cd"`

	// Type of charge in a proprietary form, as defined by the issuer.
	Proprietary *GenericIdentification3 `xml:"Prtry"`
}

func (c *ChargeType2Choice) SetCode(value string) {
	c.Code = (*ChargeType1Code)(&value)
}

func (c *ChargeType2Choice) AddProprietary() *GenericIdentification3 {
	c.Proprietary = new(GenericIdentification3)
	return c.Proprietary
}
