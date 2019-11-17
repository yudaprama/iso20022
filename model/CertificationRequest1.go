package model

// Certification request PKCS#10 (Public Key Certificate Standard 10) for creation or renewal of an X.509 certificate.
type CertificationRequest1 struct {

	// Information of the certificate to create.
	CertificateRequestInformation *CertificationRequest2 `xml:"CertReqInf"`

	// Identification of the key.
	KeyIdentification *Max140Text `xml:"KeyId,omitempty"`

	// Version of the key.
	KeyVersion *Max140Text `xml:"KeyVrsn,omitempty"`
}

func (c *CertificationRequest1) AddCertificateRequestInformation() *CertificationRequest2 {
	c.CertificateRequestInformation = new(CertificationRequest2)
	return c.CertificateRequestInformation
}

func (c *CertificationRequest1) SetKeyIdentification(value string) {
	c.KeyIdentification = (*Max140Text)(&value)
}

func (c *CertificationRequest1) SetKeyVersion(value string) {
	c.KeyVersion = (*Max140Text)(&value)
}
