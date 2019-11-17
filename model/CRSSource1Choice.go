package model

// Choice of formats for the source of the Common Reporting Standard (CRS).
type CRSSource1Choice struct {

	// Source of the Common Reporting Standard (CRS) status expressed as a code.
	Code *CRSSourceStatus1Code `xml:"Cd"`

	// Source of Common Reporting Standard (CRS) status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CRSSource1Choice) SetCode(value string) {
	c.Code = (*CRSSourceStatus1Code)(&value)
}

func (c *CRSSource1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
