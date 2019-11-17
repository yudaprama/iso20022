package model

// Choice of formats for the type of cash account.
type CashAccountType3Choice struct {

	// Type of cash account expressed as a code.
	Code *CashAccountType5Code `xml:"Cd"`

	// Type of cash account expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CashAccountType3Choice) SetCode(value string) {
	c.Code = (*CashAccountType5Code)(&value)
}

func (c *CashAccountType3Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
