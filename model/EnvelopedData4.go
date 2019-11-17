package model

// Encrypted data with encryption key.
type EnvelopedData4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Session key or identification of the protection key used by the recipient.
	Recipient []*Recipient4Choice `xml:"Rcpt"`

	// Data protection by encryption (digital envelope), with an encryption key.
	EncryptedContent *EncryptedContent3 `xml:"NcrptdCntt,omitempty"`
}

func (e *EnvelopedData4) SetVersion(value string) {
	e.Version = (*Number)(&value)
}

func (e *EnvelopedData4) AddRecipient() *Recipient4Choice {
	newValue := new(Recipient4Choice)
	e.Recipient = append(e.Recipient, newValue)
	return newValue
}

func (e *EnvelopedData4) AddEncryptedContent() *EncryptedContent3 {
	e.EncryptedContent = new(EncryptedContent3)
	return e.EncryptedContent
}
