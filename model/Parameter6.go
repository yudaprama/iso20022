package model

// Parameters associated to a cryptographic encryption algorithm.
type Parameter6 struct {

	// Format of data before encryption, if the format is not plaintext or implicit.
	EncryptionFormat *EncryptionFormat1Code `xml:"NcrptnFrmt,omitempty"`

	// Initialisation vector of a cipher block chaining (CBC) mode encryption.
	InitialisationVector *Max500Binary `xml:"InitlstnVctr,omitempty"`

	// Byte padding for a cypher block chaining mode encryption, if the padding is not implicit.
	BytePadding *BytePadding1Code `xml:"BPddg,omitempty"`
}

func (p *Parameter6) SetEncryptionFormat(value string) {
	p.EncryptionFormat = (*EncryptionFormat1Code)(&value)
}

func (p *Parameter6) SetInitialisationVector(value string) {
	p.InitialisationVector = (*Max500Binary)(&value)
}

func (p *Parameter6) SetBytePadding(value string) {
	p.BytePadding = (*BytePadding1Code)(&value)
}
