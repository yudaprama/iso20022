package model

// Parameters associated to a cryptographic algorithm.
type Parameter1 struct {

	// Initialisation vector of a cipher block chaining (CBC) mode encryption.
	InitialisationVector *Max500Binary `xml:"InitlstnVctr,omitempty"`
}

func (p *Parameter1) SetInitialisationVector(value string) {
	p.InitialisationVector = (*Max500Binary)(&value)
}
