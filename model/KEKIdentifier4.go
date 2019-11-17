package model

// Cryptographic key involved in the security command.
type KEKIdentifier4 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Identification of the cryptographic key.
	KeyIdentification *Max140Text `xml:"KeyId,omitempty"`

	// Version of the cryptographic key.
	KeyVersion *Max140Text `xml:"KeyVrsn,omitempty"`
}

func (k *KEKIdentifier4) SetName(value string) {
	k.Name = (*Max140Text)(&value)
}

func (k *KEKIdentifier4) SetKeyIdentification(value string) {
	k.KeyIdentification = (*Max140Text)(&value)
}

func (k *KEKIdentifier4) SetKeyVersion(value string) {
	k.KeyVersion = (*Max140Text)(&value)
}
