package model

// Cryptographic algorithm and parameters for digests.
type AlgorithmIdentification5 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm5Code `xml:"Algo"`
}

func (a *AlgorithmIdentification5) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm5Code)(&value)
}
