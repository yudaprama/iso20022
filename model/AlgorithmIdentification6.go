package model

// Cryptographic algorithm and parameters for encryptions with a symmetric cryptographic key.
type AlgorithmIdentification6 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm6Code `xml:"Algo"`

	// Parameters associated with the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification6) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm6Code)(&value)
}

func (a *AlgorithmIdentification6) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}
