package model

// Choice of formats for the specification of the charge payment method.
type ChargePaymentMethod1Choice struct {

	// Charge payment method expressed as a code.
	Code *ChargePaymentMethod1Code `xml:"Cd"`

	// Charge payment method expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ChargePaymentMethod1Choice) SetCode(value string) {
	c.Code = (*ChargePaymentMethod1Code)(&value)
}

func (c *ChargePaymentMethod1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
