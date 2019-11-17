package model

// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier5 struct {

	// Identification of the cryptographic key.
	KeyIdentification *Max140Text `xml:"KeyId"`

	// Version of the cryptographic key.
	KeyVersion *Max140Text `xml:"KeyVrsn"`

	// Number of usages of the cryptographic key.
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Identification used for derivation of a unique key from a master key provided for the data protection.
	DerivationIdentification *Min5Max16Binary `xml:"DerivtnId,omitempty"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp,omitempty"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn,omitempty"`
}

func (k *KEKIdentifier5) SetKeyIdentification(value string) {
	k.KeyIdentification = (*Max140Text)(&value)
}

func (k *KEKIdentifier5) SetKeyVersion(value string) {
	k.KeyVersion = (*Max140Text)(&value)
}

func (k *KEKIdentifier5) SetSequenceNumber(value string) {
	k.SequenceNumber = (*Number)(&value)
}

func (k *KEKIdentifier5) SetDerivationIdentification(value string) {
	k.DerivationIdentification = (*Min5Max16Binary)(&value)
}

func (k *KEKIdentifier5) SetType(value string) {
	k.Type = (*CryptographicKeyType3Code)(&value)
}

func (k *KEKIdentifier5) AddFunction(value string) {
	k.Function = append(k.Function, (*KeyUsage1Code)(&value))
}
