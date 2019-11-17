package model

// Choice of formats for the type of Common Reporting Standard (CRS) form.
type CRSForm1Choice struct {

	// Type of Foreign Account Tax Compliance Act (FATCA) form expressed as a code.
	Code *CRSFormType1Code `xml:"Cd"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CRSForm1Choice) SetCode(value string) {
	c.Code = (*CRSFormType1Code)(&value)
}

func (c *CRSForm1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
