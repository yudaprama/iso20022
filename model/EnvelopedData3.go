package model

// Encrypted data with encryption key.
type EnvelopedData3 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Transport key or key encryption key (KEK) identification for the recipient.
	Recipient []*Recipient3Choice `xml:"Rcpt"`

	// Encrypted data with an encryption key.
	EncryptedContent *EncryptedContent2 `xml:"NcrptdCntt"`
}

func (e *EnvelopedData3) SetVersion(value string) {
	e.Version = (*Number)(&value)
}

func (e *EnvelopedData3) AddRecipient() *Recipient3Choice {
	newValue := new(Recipient3Choice)
	e.Recipient = append(e.Recipient, newValue)
	return newValue
}

func (e *EnvelopedData3) AddEncryptedContent() *EncryptedContent2 {
	e.EncryptedContent = new(EncryptedContent2)
	return e.EncryptedContent
}
