package model

// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification9 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm9Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification9) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm9Code)(&value)
}

func (a *AlgorithmIdentification9) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}
