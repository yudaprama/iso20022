package model

// Encrypted data with an encryption key.
type EncryptedContent2 struct {

	// Type of data which have been encrypted.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Algorithm used to encrypt the data.
	ContentEncryptionAlgorithm *AlgorithmIdentification6 `xml:"CnttNcrptnAlgo"`

	// Encrypted data, result of the content encryption.
	EncryptedData *Max10000Binary `xml:"NcrptdData"`
}

func (e *EncryptedContent2) SetContentType(value string) {
	e.ContentType = (*ContentType1Code)(&value)
}

func (e *EncryptedContent2) AddContentEncryptionAlgorithm() *AlgorithmIdentification6 {
	e.ContentEncryptionAlgorithm = new(AlgorithmIdentification6)
	return e.ContentEncryptionAlgorithm
}

func (e *EncryptedContent2) SetEncryptedData(value string) {
	e.EncryptedData = (*Max10000Binary)(&value)
}
