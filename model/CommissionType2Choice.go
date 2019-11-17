package model

// Choice between a code or a data source scheme to determine the commission type.
type CommissionType2Choice struct {

	// Commission type is identified using a code.
	Code *CommissionType9Code `xml:"Cd"`

	// Commission type expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (c *CommissionType2Choice) SetCode(value string) {
	c.Code = (*CommissionType9Code)(&value)
}

func (c *CommissionType2Choice) AddProprietary() *GenericIdentification38 {
	c.Proprietary = new(GenericIdentification38)
	return c.Proprietary
}
