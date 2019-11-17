package model

// Information of the certificate to create.
type CertificationRequest2 struct {

	// Version of the certificate request information data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Distinguished name of the certificate subject, the entity whose public key is to be certified.
	SubjectName *CertificateIssuer1 `xml:"SbjtNm,omitempty"`

	// Information about the public key being certified.
	SubjectPublicKeyInformation *PublicRSAKey2 `xml:"SbjtPblcKeyInf"`

	// Attribute of the certificate service to be put in the certificate extensions, or to be used for the request.
	Attribute []*RelativeDistinguishedName2 `xml:"Attr"`
}

func (c *CertificationRequest2) SetVersion(value string) {
	c.Version = (*Number)(&value)
}

func (c *CertificationRequest2) AddSubjectName() *CertificateIssuer1 {
	c.SubjectName = new(CertificateIssuer1)
	return c.SubjectName
}

func (c *CertificationRequest2) AddSubjectPublicKeyInformation() *PublicRSAKey2 {
	c.SubjectPublicKeyInformation = new(PublicRSAKey2)
	return c.SubjectPublicKeyInformation
}

func (c *CertificationRequest2) AddAttribute() *RelativeDistinguishedName2 {
	newValue := new(RelativeDistinguishedName2)
	c.Attribute = append(c.Attribute, newValue)
	return newValue
}
