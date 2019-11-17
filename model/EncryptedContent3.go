package model

// Encrypted data with an encryption key.
type EncryptedContent3 struct {

	// Type of data which have been encrypted.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Algorithm used to encrypt the data.
	ContentEncryptionAlgorithm *AlgorithmIdentification14 `xml:"CnttNcrptnAlgo"`

	// Encrypted data, result of the content encryption.
	EncryptedData *Max100KBinary `xml:"NcrptdData"`
}

func (e *EncryptedContent3) SetContentType(value string) {
	e.ContentType = (*ContentType2Code)(&value)
}

func (e *EncryptedContent3) AddContentEncryptionAlgorithm() *AlgorithmIdentification14 {
	e.ContentEncryptionAlgorithm = new(AlgorithmIdentification14)
	return e.ContentEncryptionAlgorithm
}

func (e *EncryptedContent3) SetEncryptedData(value string) {
	e.EncryptedData = (*Max100KBinary)(&value)
}
