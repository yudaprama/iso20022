package model

// Reference of the transaction, that is the underlying payment instruction or statement entry.
type CertificateReference1 struct {

	// Unique identification of the underlying payment instruction or statement entry.
	Identification *CertificateIdentification1 `xml:"Id"`

	// Date of the underlying payment instruction or statement entry.
	Date *ISODate `xml:"Dt"`
}

func (c *CertificateReference1) AddIdentification() *CertificateIdentification1 {
	c.Identification = new(CertificateIdentification1)
	return c.Identification
}

func (c *CertificateReference1) SetDate(value string) {
	c.Date = (*ISODate)(&value)
}
