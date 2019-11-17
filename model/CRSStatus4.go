package model

// Common Reporting Standard (CRS) status information.
type CRSStatus4 struct {

	// Common Reporting Standard (CRS) status.
	Type *CRSStatus3Choice `xml:"Tp"`

	// Source of the Common Reporting Standard (CRS) status.
	Source *CRSSource1Choice `xml:"Src,omitempty"`

	// Reporting country for the CRS status when there is an exception at the country level.
	ExceptionalReportingCountry *CountryCode `xml:"XcptnlRptgCtry,omitempty"`
}

func (c *CRSStatus4) AddType() *CRSStatus3Choice {
	c.Type = new(CRSStatus3Choice)
	return c.Type
}

func (c *CRSStatus4) AddSource() *CRSSource1Choice {
	c.Source = new(CRSSource1Choice)
	return c.Source
}

func (c *CRSStatus4) SetExceptionalReportingCountry(value string) {
	c.ExceptionalReportingCountry = (*CountryCode)(&value)
}
