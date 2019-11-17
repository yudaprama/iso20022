package model

// Certificate issuer name and serial number  (see X.509).
type IssuerAndSerialNumber1 struct {

	// Certificate issuer name (see X.509).
	Issuer *CertificateIssuer1 `xml:"Issr"`

	// Certificate serial number (see X.509).
	SerialNumber *Max35Binary `xml:"SrlNb"`
}

func (i *IssuerAndSerialNumber1) AddIssuer() *CertificateIssuer1 {
	i.Issuer = new(CertificateIssuer1)
	return i.Issuer
}

func (i *IssuerAndSerialNumber1) SetSerialNumber(value string) {
	i.SerialNumber = (*Max35Binary)(&value)
}
