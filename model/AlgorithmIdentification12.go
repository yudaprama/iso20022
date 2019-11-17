package model

// Mask generator function cryptographic algorithm and parameters.
type AlgorithmIdentification12 struct {

	// Mask generator function cryptographic algorithm.
	Algorithm *Algorithm8Code `xml:"Algo"`

	// Parameters associated to the mask generator function cryptographic algorithm
	Parameter *Parameter5 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification12) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm8Code)(&value)
}

func (a *AlgorithmIdentification12) AddParameter() *Parameter5 {
	a.Parameter = new(Parameter5)
	return a.Parameter
}
