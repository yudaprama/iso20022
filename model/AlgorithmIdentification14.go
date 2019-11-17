package model

// Cryptographic algorithm and parameters for encryptions with a symmetric cryptographic key.
type AlgorithmIdentification14 struct {

	// Identification of the encryption algorithm.
	Algorithm *Algorithm15Code `xml:"Algo"`

	// Parameters associated with the CBC (Chain Block Chaining) encryption algorithm.
	Parameter *Parameter6 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification14) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm15Code)(&value)
}

func (a *AlgorithmIdentification14) AddParameter() *Parameter6 {
	a.Parameter = new(Parameter6)
	return a.Parameter
}
