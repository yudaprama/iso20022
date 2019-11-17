package model

// Choice of formats for the civil status.
type CivilStatus1Choice struct {

	// Civil status expressed as a code.
	Code *CivilStatus1Code `xml:"Cd"`

	// Civil status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CivilStatus1Choice) SetCode(value string) {
	c.Code = (*CivilStatus1Code)(&value)
}

func (c *CivilStatus1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
