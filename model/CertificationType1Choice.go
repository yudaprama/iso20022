package model

// Choice of formats for the specification of the certificate type.
type CertificationType1Choice struct {

	// Certificate type expressed as a code.
	Code *CertificateType2Code `xml:"Cd"`

	// Certificate type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CertificationType1Choice) SetCode(value string) {
	c.Code = (*CertificateType2Code)(&value)
}

func (c *CertificationType1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
