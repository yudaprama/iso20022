package model

// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType5 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData2 `xml:"EnvlpdData"`
}

func (c *ContentInformationType5) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType5) AddEnvelopedData() *EnvelopedData2 {
	c.EnvelopedData = new(EnvelopedData2)
	return c.EnvelopedData
}
