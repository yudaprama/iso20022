package model

// Encrypted data with encryption key.
type EnvelopedData1 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Transport key or key encryption key (KEK) identification for the recipient.
	Recipient []*Recipient1Choice `xml:"Rcpt"`

	// Encrypted data with an encryption key.
	EncryptedContent *EncryptedContent1 `xml:"NcrptdCntt"`
}

func (e *EnvelopedData1) SetVersion(value string) {
	e.Version = (*Number)(&value)
}

func (e *EnvelopedData1) AddRecipient() *Recipient1Choice {
	newValue := new(Recipient1Choice)
	e.Recipient = append(e.Recipient, newValue)
	return newValue
}

func (e *EnvelopedData1) AddEncryptedContent() *EncryptedContent1 {
	e.EncryptedContent = new(EncryptedContent1)
	return e.EncryptedContent
}
