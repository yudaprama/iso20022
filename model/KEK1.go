package model

// Key encryption key (KEK), using previously distributed symmetric key.
type KEK1 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the key encryption key (KEK).
	KEKIdentification *KEKIdentifier1 `xml:"KEKId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification1 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KEK1) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KEK1) AddKEKIdentification() *KEKIdentifier1 {
	k.KEKIdentification = new(KEKIdentifier1)
	return k.KEKIdentification
}

func (k *KEK1) AddKeyEncryptionAlgorithm() *AlgorithmIdentification1 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification1)
	return k.KeyEncryptionAlgorithm
}

func (k *KEK1) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}
