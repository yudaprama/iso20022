package model

// Choice of format for the legal jurisdiction.
type CountrySubdivision1Choice struct {

	// Country subdivision of jurisdiction.
	//
	Code *Max35Text `xml:"Cd"`

	// Jurisdiction expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (c *CountrySubdivision1Choice) SetCode(value string) {
	c.Code = (*Max35Text)(&value)
}

func (c *CountrySubdivision1Choice) AddProprietary() *GenericIdentification1 {
	c.Proprietary = new(GenericIdentification1)
	return c.Proprietary
}
