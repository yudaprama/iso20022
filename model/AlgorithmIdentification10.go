package model

// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification10 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm10Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification10) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm10Code)(&value)
}

func (a *AlgorithmIdentification10) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}
