package model

// Choice of formats for the commission type.
type CommissionType5Choice struct {

	// Commission type expressed as a code.
	Code *CommissionType6Code `xml:"Cd"`

	// Commission type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CommissionType5Choice) SetCode(value string) {
	c.Code = (*CommissionType6Code)(&value)
}

func (c *CommissionType5Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
