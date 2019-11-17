package model

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType15 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData4 `xml:"AuthntcdData"`
}

func (c *ContentInformationType15) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType15) AddAuthenticatedData() *AuthenticatedData4 {
	c.AuthenticatedData = new(AuthenticatedData4)
	return c.AuthenticatedData
}
