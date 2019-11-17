package model

// Defines a cryptographic digest algorithm and value.
type AlgorithmAndDigest1 struct {

	// Digest algorithm used to create the digest.
	DigestAlgorithm *Algorithm5Code `xml:"DgstAlgo"`

	// Result of data-digesting process.
	Digest *Max140Text `xml:"Dgst"`
}

func (a *AlgorithmAndDigest1) SetDigestAlgorithm(value string) {
	a.DigestAlgorithm = (*Algorithm5Code)(&value)
}

func (a *AlgorithmAndDigest1) SetDigest(value string) {
	a.Digest = (*Max140Text)(&value)
}
