package model

// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification2 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm2Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification2) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm2Code)(&value)
}

func (a *AlgorithmIdentification2) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}
