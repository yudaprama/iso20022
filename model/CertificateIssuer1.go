package model

// Certificate issuer name (see X.509).
type CertificateIssuer1 struct {

	// Relative distinguished name inside a X.509 certificate.
	RelativeDistinguishedName []*RelativeDistinguishedName1 `xml:"RltvDstngshdNm"`
}

func (c *CertificateIssuer1) AddRelativeDistinguishedName() *RelativeDistinguishedName1 {
	newValue := new(RelativeDistinguishedName1)
	c.RelativeDistinguishedName = append(c.RelativeDistinguishedName, newValue)
	return newValue
}
