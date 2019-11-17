package model

// Choice of formats for the Common Reporting Standard (CRS) status.
type CRSStatus3Choice struct {

	// Common Reporting Standard (CRS) status expressed as a code.
	Code *CRSStatus1Code `xml:"Cd"`

	// Common Reporting Standard (CRS) status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CRSStatus3Choice) SetCode(value string) {
	c.Code = (*CRSStatus1Code)(&value)
}

func (c *CRSStatus3Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
