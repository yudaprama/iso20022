package model

// Mask generator function cryptographic algorithm and parameters.
type AlgorithmIdentification8 struct {

	// Mask generator function cryptographic algorithm.
	Algorithm *Algorithm8Code `xml:"Algo"`

	// Parameters associated to the mask generator function cryptographic algorithm
	Parameter *Parameter3 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification8) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm8Code)(&value)
}

func (a *AlgorithmIdentification8) AddParameter() *Parameter3 {
	a.Parameter = new(Parameter3)
	return a.Parameter
}
