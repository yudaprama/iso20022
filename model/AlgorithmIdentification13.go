package model

// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification13 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm13Code `xml:"Algo"`

	// Parameters associated to the encryption algorithm.
	Parameter *Parameter6 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification13) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm13Code)(&value)
}

func (a *AlgorithmIdentification13) AddParameter() *Parameter6 {
	a.Parameter = new(Parameter6)
	return a.Parameter
}
