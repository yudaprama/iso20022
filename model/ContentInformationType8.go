package model

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType8 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData []*AuthenticatedData3 `xml:"AuthntcdData,omitempty"`
}

func (c *ContentInformationType8) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType8) AddAuthenticatedData() *AuthenticatedData3 {
	newValue := new(AuthenticatedData3)
	c.AuthenticatedData = append(c.AuthenticatedData, newValue)
	return newValue
}
