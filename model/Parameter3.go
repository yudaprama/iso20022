package model

// Parameters associated to a mask generator cryptographic function.
type Parameter3 struct {

	// Digest algorithm used in the mask generator function.
	DigestAlgorithm *Algorithm5Code `xml:"DgstAlgo,omitempty"`
}

func (p *Parameter3) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm5Code)(&value)
}
