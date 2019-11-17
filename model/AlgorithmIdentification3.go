package model

// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification3 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm3Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification3) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm3Code)(&value)
}

func (a *AlgorithmIdentification3) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}
