package model

// Cryptographic key.
type CryptographicKey8 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id,omitempty"`

	// Identification of the security domain.
	SecurityDomainIdentification *Max35Text `xml:"SctyDomnId,omitempty"`

	// Additional identification of the key, for instance to derive the key.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Sequence counter of the cryptographic key.
	SequenceCounter *Number `xml:"SeqCntr,omitempty"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Value for checking a cryptographic key.
	KeyCheckValue *Max35Binary `xml:"KeyChckVal,omitempty"`

	// Encrypted value of the cryptographic key.
	EncryptedKeyValue *ContentInformationType10 `xml:"NcrptdKeyVal,omitempty"`

	// Value of the public component of a RSA key.
	PublicKeyValue *PublicRSAKey1 `xml:"PblcKeyVal,omitempty"`
}

func (c *CryptographicKey8) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *CryptographicKey8) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey8) SetSecurityDomainIdentification(value string) {
	c.SecurityDomainIdentification = (*Max35Text)(&value)
}

func (c *CryptographicKey8) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey8) SetVersion(value string) {
	c.Version = (*Max256Text)(&value)
}

func (c *CryptographicKey8) SetSequenceCounter(value string) {
	c.SequenceCounter = (*Number)(&value)
}

func (c *CryptographicKey8) SetType(value string) {
	c.Type = (*CryptographicKeyType3Code)(&value)
}

func (c *CryptographicKey8) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey8) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey8) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey8) SetKeyCheckValue(value string) {
	c.KeyCheckValue = (*Max35Binary)(&value)
}

func (c *CryptographicKey8) AddEncryptedKeyValue() *ContentInformationType10 {
	c.EncryptedKeyValue = new(ContentInformationType10)
	return c.EncryptedKeyValue
}

func (c *CryptographicKey8) AddPublicKeyValue() *PublicRSAKey1 {
	c.PublicKeyValue = new(PublicRSAKey1)
	return c.PublicKeyValue
}
