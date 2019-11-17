package model

// Encrypted data with an encryption key.
type EncryptedContent1 struct {

	// Type of data which have been encrypted.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Algorithm used to encrypt the data.
	ContentEncryptionAlgorithm *AlgorithmIdentification1 `xml:"CnttNcrptnAlgo"`

	// Encrypted data, result of the content encryption.
	EncryptedData *Max10000Binary `xml:"NcrptdData"`
}

func (e *EncryptedContent1) SetContentType(value string) {
	e.ContentType = (*ContentType1Code)(&value)
}

func (e *EncryptedContent1) AddContentEncryptionAlgorithm() *AlgorithmIdentification1 {
	e.ContentEncryptionAlgorithm = new(AlgorithmIdentification1)
	return e.ContentEncryptionAlgorithm
}

func (e *EncryptedContent1) SetEncryptedData(value string) {
	e.EncryptedData = (*Max10000Binary)(&value)
}
