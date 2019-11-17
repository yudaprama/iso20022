package model

// Key encryption key (KEK), encrypted previously distributed symmetric key.
type KeyTransport1 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn"`

	// Transport key or key encryption key (KEK) for the recipient.
	RecipientIdentification *CertificateIdentifier1 `xml:"RcptId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification1 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KeyTransport1) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KeyTransport1) AddRecipientIdentification() *CertificateIdentifier1 {
	k.RecipientIdentification = new(CertificateIdentifier1)
	return k.RecipientIdentification
}

func (k *KeyTransport1) AddKeyEncryptionAlgorithm() *AlgorithmIdentification1 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification1)
	return k.KeyEncryptionAlgorithm
}

func (k *KeyTransport1) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}
