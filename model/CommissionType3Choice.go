package model

// Choice of formats for the commission type.
type CommissionType3Choice struct {

	// Type of commission expressed as a code.
	Code *CommissionType7Code `xml:"Cd"`

	// Type of commission expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CommissionType3Choice) SetCode(value string) {
	c.Code = (*CommissionType7Code)(&value)
}

func (c *CommissionType3Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
