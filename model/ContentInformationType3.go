package model

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType3 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData []*AuthenticatedData1 `xml:"AuthntcdData,omitempty"`
}

func (c *ContentInformationType3) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType3) AddAuthenticatedData() *AuthenticatedData1 {
	newValue := new(AuthenticatedData1)
	c.AuthenticatedData = append(c.AuthenticatedData, newValue)
	return newValue
}
