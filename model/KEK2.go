package model

// Key encryption key (KEK), using previously distributed symmetric key.
type KEK2 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the key encryption key (KEK).
	KEKIdentification *KEKIdentifier1 `xml:"KEKId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification2 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KEK2) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KEK2) AddKEKIdentification() *KEKIdentifier1 {
	k.KEKIdentification = new(KEKIdentifier1)
	return k.KEKIdentification
}

func (k *KEK2) AddKeyEncryptionAlgorithm() *AlgorithmIdentification2 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification2)
	return k.KeyEncryptionAlgorithm
}

func (k *KEK2) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}
