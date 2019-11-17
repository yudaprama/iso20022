package model

// Specifies the charge type.
type ChargeType3Choice struct {

	// Charge type, in a coded form.
	Code *ExternalChargeType1Code `xml:"Cd"`

	// Type of charge in a proprietary form, as defined by the issuer.
	Proprietary *GenericIdentification3 `xml:"Prtry"`
}

func (c *ChargeType3Choice) SetCode(value string) {
	c.Code = (*ExternalChargeType1Code)(&value)
}

func (c *ChargeType3Choice) AddProprietary() *GenericIdentification3 {
	c.Proprietary = new(GenericIdentification3)
	return c.Proprietary
}
