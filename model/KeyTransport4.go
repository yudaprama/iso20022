package model

// Key encryption key (KEK), encrypted with a previously distributed asymmetric public key.
type KeyTransport4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a cryptographic asymmetric key for the recipient.
	RecipientIdentification *Recipient5Choice `xml:"RcptId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification11 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max5000Binary `xml:"NcrptdKey"`
}

func (k *KeyTransport4) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KeyTransport4) AddRecipientIdentification() *Recipient5Choice {
	k.RecipientIdentification = new(Recipient5Choice)
	return k.RecipientIdentification
}

func (k *KeyTransport4) AddKeyEncryptionAlgorithm() *AlgorithmIdentification11 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification11)
	return k.KeyEncryptionAlgorithm
}

func (k *KeyTransport4) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max5000Binary)(&value)
}
