package model

// Key encryption key (KEK), using previously distributed symmetric key.
type KEK3 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the key encryption key (KEK).
	KEKIdentification *KEKIdentifier1 `xml:"KEKId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification9 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KEK3) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KEK3) AddKEKIdentification() *KEKIdentifier1 {
	k.KEKIdentification = new(KEKIdentifier1)
	return k.KEKIdentification
}

func (k *KEK3) AddKeyEncryptionAlgorithm() *AlgorithmIdentification9 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification9)
	return k.KeyEncryptionAlgorithm
}

func (k *KEK3) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}
