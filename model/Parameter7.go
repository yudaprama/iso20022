package model

// Parameters associated to the MAC algorithm.
type Parameter7 struct {

	// Initialisation vector of a cipher block chaining (CBC) mode encryption.
	InitialisationVector *Max500Binary `xml:"InitlstnVctr,omitempty"`

	// Byte padding for a cypher block chaining mode encryption, if the padding is not implicit.
	BytePadding *BytePadding1Code `xml:"BPddg,omitempty"`
}

func (p *Parameter7) SetInitialisationVector(value string) {
	p.InitialisationVector = (*Max500Binary)(&value)
}

func (p *Parameter7) SetBytePadding(value string) {
	p.BytePadding = (*BytePadding1Code)(&value)
}
