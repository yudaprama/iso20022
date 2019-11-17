package model

// Identification of a cryptographic algorithm.
type AlgorithmIdentification1 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm1Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification1) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm1Code)(&value)
}

func (a *AlgorithmIdentification1) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}
