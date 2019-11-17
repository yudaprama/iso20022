package model

// Cryptographic Key.
type CryptographicKey6 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Version of the cryptographic key.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Encrypted value of the key present as CMS structure EnvelopedData.
	EncryptedKeyValue *ContentInformationType10 `xml:"NcrptdKeyVal,omitempty"`

	// Certificate to protect the key.
	Certificate []*Max5000Binary `xml:"Cert,omitempty"`

	// Chip card protection of the key.
	ICCRelatedData *Max5000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CryptographicKey6) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *CryptographicKey6) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey6) SetVersion(value string) {
	c.Version = (*Max256Text)(&value)
}

func (c *CryptographicKey6) SetType(value string) {
	c.Type = (*CryptographicKeyType3Code)(&value)
}

func (c *CryptographicKey6) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey6) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey6) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey6) AddEncryptedKeyValue() *ContentInformationType10 {
	c.EncryptedKeyValue = new(ContentInformationType10)
	return c.EncryptedKeyValue
}

func (c *CryptographicKey6) AddCertificate(value string) {
	c.Certificate = append(c.Certificate, (*Max5000Binary)(&value))
}

func (c *CryptographicKey6) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max5000Binary)(&value)
}
