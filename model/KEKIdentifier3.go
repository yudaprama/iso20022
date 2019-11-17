package model

// Key that must be created and sent in the response, or that must be verified..
type KEKIdentifier3 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Identification of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Version of the cryptographic key.
	Version *Max140Text `xml:"Vrsn,omitempty"`

	// Value to check the key, for instance, result of the encryption of the null binary string.
	KeyCheckValue *Max35Binary `xml:"KeyChckVal,omitempty"`
}

func (k *KEKIdentifier3) SetName(value string) {
	k.Name = (*Max140Text)(&value)
}

func (k *KEKIdentifier3) SetIdentification(value string) {
	k.Identification = (*Max140Text)(&value)
}

func (k *KEKIdentifier3) SetVersion(value string) {
	k.Version = (*Max140Text)(&value)
}

func (k *KEKIdentifier3) SetKeyCheckValue(value string) {
	k.KeyCheckValue = (*Max35Binary)(&value)
}
