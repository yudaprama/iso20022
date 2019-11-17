package model

// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification15 struct {

	// Identification of the MAC algorithm.
	Algorithm *Algorithm12Code `xml:"Algo"`

	// Parameters associated to the MAC algorithm.
	Parameter *Parameter7 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification15) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm12Code)(&value)
}

func (a *AlgorithmIdentification15) AddParameter() *Parameter7 {
	a.Parameter = new(Parameter7)
	return a.Parameter
}
