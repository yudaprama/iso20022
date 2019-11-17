package model

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType6 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData []*AuthenticatedData2 `xml:"AuthntcdData,omitempty"`
}

func (c *ContentInformationType6) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType6) AddAuthenticatedData() *AuthenticatedData2 {
	newValue := new(AuthenticatedData2)
	c.AuthenticatedData = append(c.AuthenticatedData, newValue)
	return newValue
}
