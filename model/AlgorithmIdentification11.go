package model

// Cryptographic algorithms and parameters for the protection of transported keys by an asymmetric key.
type AlgorithmIdentification11 struct {

	// Asymmetric encryption algorithm of a transport key.
	Algorithm *Algorithm7Code `xml:"Algo"`

	// Parameters of the encryption algorithm.
	Parameter *Parameter4 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification11) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm7Code)(&value)
}

func (a *AlgorithmIdentification11) AddParameter() *Parameter4 {
	a.Parameter = new(Parameter4)
	return a.Parameter
}
