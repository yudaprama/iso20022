package model

// General cryptographic message syntax (CMS) containing protected data.
type ContentInformationType1 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData1 `xml:"EnvlpdData,omitempty"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData1 `xml:"AuthntcdData,omitempty"`

	// Data protected by digital signatures.
	SignedData *SignedData1 `xml:"SgndData,omitempty"`

	// Data protected by a digest.
	DigestedData *DigestedData1 `xml:"DgstdData,omitempty"`

	// Data protection by encryption with a previously exchanged key identified by a name.
	NamedKeyEncryptedData *NamedKeyEncryptedData1 `xml:"NmdKeyNcrptdData,omitempty"`
}

func (c *ContentInformationType1) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType1) AddEnvelopedData() *EnvelopedData1 {
	c.EnvelopedData = new(EnvelopedData1)
	return c.EnvelopedData
}

func (c *ContentInformationType1) AddAuthenticatedData() *AuthenticatedData1 {
	c.AuthenticatedData = new(AuthenticatedData1)
	return c.AuthenticatedData
}

func (c *ContentInformationType1) AddSignedData() *SignedData1 {
	c.SignedData = new(SignedData1)
	return c.SignedData
}

func (c *ContentInformationType1) AddDigestedData() *DigestedData1 {
	c.DigestedData = new(DigestedData1)
	return c.DigestedData
}

func (c *ContentInformationType1) AddNamedKeyEncryptedData() *NamedKeyEncryptedData1 {
	c.NamedKeyEncryptedData = new(NamedKeyEncryptedData1)
	return c.NamedKeyEncryptedData
}
