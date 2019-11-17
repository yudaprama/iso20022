package model

// Encrypted data with an encryption key identified with a name.
type NamedKeyEncryptedData1 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Name of the key encryption key (KEK).
	KeyName *Max140Text `xml:"KeyNm,omitempty"`

	// Encrypted data with an encryption key.
	EncryptedContent *EncryptedContent1 `xml:"NcrptdCntt"`
}

func (n *NamedKeyEncryptedData1) SetVersion(value string) {
	n.Version = (*Number)(&value)
}

func (n *NamedKeyEncryptedData1) SetKeyName(value string) {
	n.KeyName = (*Max140Text)(&value)
}

func (n *NamedKeyEncryptedData1) AddEncryptedContent() *EncryptedContent1 {
	n.EncryptedContent = new(EncryptedContent1)
	return n.EncryptedContent
}
