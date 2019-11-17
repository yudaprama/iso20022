package model

// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier1 struct {

	// Identification of the cryptographic key.
	KeyIdentification *Max140Text `xml:"KeyId"`

	// Version of the cryptographic key.
	KeyVersion *Exact10Text `xml:"KeyVrsn"`

	// Identification used for derivation of a unique key from a master key provided for the data protection.
	DerivationIdentification *Min5Max16Binary `xml:"DerivtnId,omitempty"`
}

func (k *KEKIdentifier1) SetKeyIdentification(value string) {
	k.KeyIdentification = (*Max140Text)(&value)
}

func (k *KEKIdentifier1) SetKeyVersion(value string) {
	k.KeyVersion = (*Exact10Text)(&value)
}

func (k *KEKIdentifier1) SetDerivationIdentification(value string) {
	k.DerivationIdentification = (*Min5Max16Binary)(&value)
}
