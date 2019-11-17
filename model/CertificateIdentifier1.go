package model

// Transport key or key encryption key (KEK) identification for the recipient.
type CertificateIdentifier1 struct {

	// Certificate issuer name and serial number  (see X.509).
	IssuerAndSerialNumber *IssuerAndSerialNumber1 `xml:"IssrAndSrlNb"`
}

func (c *CertificateIdentifier1) AddIssuerAndSerialNumber() *IssuerAndSerialNumber1 {
	c.IssuerAndSerialNumber = new(IssuerAndSerialNumber1)
	return c.IssuerAndSerialNumber
}
