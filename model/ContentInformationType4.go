package model

// General cryptographic message syntax (CMS) containing protected data.
type ContentInformationType4 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData2 `xml:"EnvlpdData,omitempty"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData2 `xml:"AuthntcdData,omitempty"`

	// Data protected by digital signatures.
	SignedData *SignedData2 `xml:"SgndData,omitempty"`

	// Data protected by a digest.
	DigestedData *DigestedData2 `xml:"DgstdData,omitempty"`

	// Data protection by encryption with a previously exchanged key identified by a name.
	NamedKeyEncryptedData *NamedKeyEncryptedData2 `xml:"NmdKeyNcrptdData,omitempty"`
}

func (c *ContentInformationType4) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType4) AddEnvelopedData() *EnvelopedData2 {
	c.EnvelopedData = new(EnvelopedData2)
	return c.EnvelopedData
}

func (c *ContentInformationType4) AddAuthenticatedData() *AuthenticatedData2 {
	c.AuthenticatedData = new(AuthenticatedData2)
	return c.AuthenticatedData
}

func (c *ContentInformationType4) AddSignedData() *SignedData2 {
	c.SignedData = new(SignedData2)
	return c.SignedData
}

func (c *ContentInformationType4) AddDigestedData() *DigestedData2 {
	c.DigestedData = new(DigestedData2)
	return c.DigestedData
}

func (c *ContentInformationType4) AddNamedKeyEncryptedData() *NamedKeyEncryptedData2 {
	c.NamedKeyEncryptedData = new(NamedKeyEncryptedData2)
	return c.NamedKeyEncryptedData
}
