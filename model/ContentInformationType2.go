package model

// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType2 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData1 `xml:"EnvlpdData"`
}

func (c *ContentInformationType2) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType2) AddEnvelopedData() *EnvelopedData1 {
	c.EnvelopedData = new(EnvelopedData1)
	return c.EnvelopedData
}
