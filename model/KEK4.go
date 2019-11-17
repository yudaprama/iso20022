package model

// Key encryption key (KEK), using previously distributed symmetric key.
type KEK4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the key encryption key (KEK).
	KEKIdentification *KEKIdentifier2 `xml:"KEKId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification13 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max500Binary `xml:"NcrptdKey"`
}

func (k *KEK4) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KEK4) AddKEKIdentification() *KEKIdentifier2 {
	k.KEKIdentification = new(KEKIdentifier2)
	return k.KEKIdentification
}

func (k *KEK4) AddKeyEncryptionAlgorithm() *AlgorithmIdentification13 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification13)
	return k.KeyEncryptionAlgorithm
}

func (k *KEK4) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max500Binary)(&value)
}
