package model

// Cryptographic algorithm and parameters of digests.
type AlgorithmIdentification16 struct {

	// Identification of the digest algorithm.
	Algorithm *Algorithm11Code `xml:"Algo"`
}

func (a *AlgorithmIdentification16) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm11Code)(&value)
}
