package model

// Parameters associated to a mask generator cryptographic function.
type Parameter5 struct {

	// Digest algorithm used in the mask generator function.
	DigestAlgorithm *Algorithm11Code `xml:"DgstAlgo,omitempty"`
}

func (p *Parameter5) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm11Code)(&value)
}
