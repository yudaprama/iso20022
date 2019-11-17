package model

// Key encryption key (KEK), encrypted with a previously distributed asymmetric public key.
type KeyTransport2 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn"`

	// Transport key or key encryption key (KEK) for the recipient.
	RecipientIdentification *CertificateIdentifier1 `xml:"RcptId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification7 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KeyTransport2) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KeyTransport2) AddRecipientIdentification() *CertificateIdentifier1 {
	k.RecipientIdentification = new(CertificateIdentifier1)
	return k.RecipientIdentification
}

func (k *KeyTransport2) AddKeyEncryptionAlgorithm() *AlgorithmIdentification7 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification7)
	return k.KeyEncryptionAlgorithm
}

func (k *KeyTransport2) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}
