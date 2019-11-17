package model

// General cryptographic message syntax (CMS) containing protected data.
type ContentInformationType9 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData3 `xml:"EnvlpdData,omitempty"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData3 `xml:"AuthntcdData,omitempty"`

	// Data protected by digital signatures.
	SignedData *SignedData3 `xml:"SgndData,omitempty"`

	// Data protected by a digest.
	DigestedData *DigestedData3 `xml:"DgstdData,omitempty"`
}

func (c *ContentInformationType9) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType9) AddEnvelopedData() *EnvelopedData3 {
	c.EnvelopedData = new(EnvelopedData3)
	return c.EnvelopedData
}

func (c *ContentInformationType9) AddAuthenticatedData() *AuthenticatedData3 {
	c.AuthenticatedData = new(AuthenticatedData3)
	return c.AuthenticatedData
}

func (c *ContentInformationType9) AddSignedData() *SignedData3 {
	c.SignedData = new(SignedData3)
	return c.SignedData
}

func (c *ContentInformationType9) AddDigestedData() *DigestedData3 {
	c.DigestedData = new(DigestedData3)
	return c.DigestedData
}
