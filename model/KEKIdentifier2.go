package model

// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier2 struct {

	// Identification of the cryptographic key.
	KeyIdentification *Max140Text `xml:"KeyId"`

	// Version of the cryptographic key.
	KeyVersion *Max140Text `xml:"KeyVrsn"`

	// Number of usages of the cryptographic key.
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Identification used for derivation of a unique key from a master key provided for the data protection.
	DerivationIdentification *Min5Max16Binary `xml:"DerivtnId,omitempty"`
}

func (k *KEKIdentifier2) SetKeyIdentification(value string) {
	k.KeyIdentification = (*Max140Text)(&value)
}

func (k *KEKIdentifier2) SetKeyVersion(value string) {
	k.KeyVersion = (*Max140Text)(&value)
}

func (k *KEKIdentifier2) SetSequenceNumber(value string) {
	k.SequenceNumber = (*Number)(&value)
}

func (k *KEKIdentifier2) SetDerivationIdentification(value string) {
	k.DerivationIdentification = (*Min5Max16Binary)(&value)
}
