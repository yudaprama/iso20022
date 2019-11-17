package model

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType11 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData []*AuthenticatedData4 `xml:"AuthntcdData,omitempty"`
}

func (c *ContentInformationType11) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType11) AddAuthenticatedData() *AuthenticatedData4 {
	newValue := new(AuthenticatedData4)
	c.AuthenticatedData = append(c.AuthenticatedData, newValue)
	return newValue
}
