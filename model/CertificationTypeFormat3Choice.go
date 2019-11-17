package model

// Choice between a standard code or proprietary code to specify the certification format required.
type CertificationTypeFormat3Choice struct {

	// Standard code to specify the certification format required, that is, physical or electronic format.
	Code *CertificationFormatType1Code `xml:"Cd"`

	// Proprietary identification of the certification format.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CertificationTypeFormat3Choice) SetCode(value string) {
	c.Code = (*CertificationFormatType1Code)(&value)
}

func (c *CertificationTypeFormat3Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
