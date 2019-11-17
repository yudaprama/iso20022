package model

// Identification of a cryptographic algorithm and parameters for digital signatures.
type AlgorithmIdentification4 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm4Code `xml:"Algo"`
}

func (a *AlgorithmIdentification4) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm4Code)(&value)
}
