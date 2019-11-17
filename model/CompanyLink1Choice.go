package model

// Choice of formats for the trading party or broker.
type CompanyLink1Choice struct {

	// Company link expressed as a code.
	Code *CompanyLink1Code `xml:"Cd"`

	// Company link expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CompanyLink1Choice) SetCode(value string) {
	c.Code = (*CompanyLink1Code)(&value)
}

func (c *CompanyLink1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
