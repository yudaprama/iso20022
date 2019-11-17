package model

// Cryptographic Key.
type CryptographicKey5 struct {

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Additional identification of the key.
	// Usage
	// For derived unique key per transaction (DUKPT) keys, the key serial number (KSN) with the 21 bits of the transaction counter set to zero.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Max256Text `xml:"Vrsn"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Encrypted cryptographic key.
	KeyValue *ContentInformationType10 `xml:"KeyVal"`
}

func (c *CryptographicKey5) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey5) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey5) SetVersion(value string) {
	c.Version = (*Max256Text)(&value)
}

func (c *CryptographicKey5) SetType(value string) {
	c.Type = (*CryptographicKeyType3Code)(&value)
}

func (c *CryptographicKey5) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey5) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey5) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey5) AddKeyValue() *ContentInformationType10 {
	c.KeyValue = new(ContentInformationType10)
	return c.KeyValue
}
