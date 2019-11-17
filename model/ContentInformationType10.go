package model

// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType10 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by encryption or by a digital envelope, with an encryption key.
	EnvelopedData *EnvelopedData4 `xml:"EnvlpdData"`
}

func (c *ContentInformationType10) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType10) AddEnvelopedData() *EnvelopedData4 {
	c.EnvelopedData = new(EnvelopedData4)
	return c.EnvelopedData
}
