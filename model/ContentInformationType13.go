package model

// General cryptographic message syntax (CMS) containing data. protected by a MAC or a digital signature
type ContentInformationType13 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData4 `xml:"AuthntcdData,omitempty"`

	// Data protected by a digital signatures.
	SignedData *SignedData4 `xml:"SgndData,omitempty"`
}

func (c *ContentInformationType13) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType13) AddAuthenticatedData() *AuthenticatedData4 {
	c.AuthenticatedData = new(AuthenticatedData4)
	return c.AuthenticatedData
}

func (c *ContentInformationType13) AddSignedData() *SignedData4 {
	c.SignedData = new(SignedData4)
	return c.SignedData
}
