package model

// General cryptographic message syntax (CMS) containing data. protected by a digital signature
type ContentInformationType14 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protected by a digital signatures.
	SignedData *SignedData4 `xml:"SgndData"`
}

func (c *ContentInformationType14) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType14) AddSignedData() *SignedData4 {
	c.SignedData = new(SignedData4)
	return c.SignedData
}
