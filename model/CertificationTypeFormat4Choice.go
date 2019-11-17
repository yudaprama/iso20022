package model

// Choice between a standard code or proprietary code to specify the certification format required.
type CertificationTypeFormat4Choice struct {

	// Standard code to specify the certification format required, that is, physical or electronic format.
	Code *CertificationFormatType1Code `xml:"Cd"`

	// Proprietary identification of the certification format.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CertificationTypeFormat4Choice) SetCode(value string) {
	c.Code = (*CertificationFormatType1Code)(&value)
}

func (c *CertificationTypeFormat4Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
