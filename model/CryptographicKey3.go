package model

// Indetification of a cryptographic Key to use.
type CryptographicKey3 struct {

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Additional identification of the key.
	// Usage
	// For derived unique key per transaction (DUKPT) keys, the key serial number (KSN) with the 21 bits of the transaction counter set to zero.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Exact10Text `xml:"Vrsn"`
}

func (c *CryptographicKey3) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey3) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey3) SetVersion(value string) {
	c.Version = (*Exact10Text)(&value)
}
