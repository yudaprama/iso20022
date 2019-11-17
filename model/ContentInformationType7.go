package model

// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType7 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData3 `xml:"EnvlpdData"`
}

func (c *ContentInformationType7) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType7) AddEnvelopedData() *EnvelopedData3 {
	c.EnvelopedData = new(EnvelopedData3)
	return c.EnvelopedData
}
