package model

// Encrypted data with encryption key.
type EnvelopedData2 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Transport key or key encryption key (KEK) identification for the recipient.
	Recipient []*Recipient2Choice `xml:"Rcpt"`

	// Encrypted data with an encryption key.
	EncryptedContent *EncryptedContent2 `xml:"NcrptdCntt"`
}

func (e *EnvelopedData2) SetVersion(value string) {
	e.Version = (*Number)(&value)
}

func (e *EnvelopedData2) AddRecipient() *Recipient2Choice {
	newValue := new(Recipient2Choice)
	e.Recipient = append(e.Recipient, newValue)
	return newValue
}

func (e *EnvelopedData2) AddEncryptedContent() *EncryptedContent2 {
	e.EncryptedContent = new(EncryptedContent2)
	return e.EncryptedContent
}
