package model

// General cryptographic message syntax (CMS) containing protected data.
type ContentInformationType12 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData4 `xml:"EnvlpdData,omitempty"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData4 `xml:"AuthntcdData,omitempty"`

	// Data protected by a digital signatures.
	SignedData *SignedData4 `xml:"SgndData,omitempty"`

	// Data protected by a digest.
	DigestedData *DigestedData4 `xml:"DgstdData,omitempty"`
}

func (c *ContentInformationType12) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType12) AddEnvelopedData() *EnvelopedData4 {
	c.EnvelopedData = new(EnvelopedData4)
	return c.EnvelopedData
}

func (c *ContentInformationType12) AddAuthenticatedData() *AuthenticatedData4 {
	c.AuthenticatedData = new(AuthenticatedData4)
	return c.AuthenticatedData
}

func (c *ContentInformationType12) AddSignedData() *SignedData4 {
	c.SignedData = new(SignedData4)
	return c.SignedData
}

func (c *ContentInformationType12) AddDigestedData() *DigestedData4 {
	c.DigestedData = new(DigestedData4)
	return c.DigestedData
}
