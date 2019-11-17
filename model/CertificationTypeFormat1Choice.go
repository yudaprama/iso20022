package model

// Choice between a standard code or proprietary code to specify the certification format required.
type CertificationTypeFormat1Choice struct {

	// Standard code to specify the certification format required, that is, physical or electronic format.
	Code *CertificationFormatType1Code `xml:"Cd"`

	// Proprietary identification of the certification format.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CertificationTypeFormat1Choice) SetCode(value string) {
	c.Code = (*CertificationFormatType1Code)(&value)
}

func (c *CertificationTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
