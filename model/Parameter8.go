package model

// Parameters of the RSASSA-PSS digital signature algorithm (RSA signature algorithm with appendix: Probabilistic Signature Scheme).
type Parameter8 struct {

	// Identification of the digest algorithm.
	DigestAlgorithm *Algorithm11Code `xml:"DgstAlgo"`

	// Mask generator function cryptographic algorithm and parameters.
	MaskGeneratorAlgorithm *AlgorithmIdentification12 `xml:"MskGnrtrAlgo"`

	// Length of the salt to include in the signature.
	SaltLength *Number `xml:"SaltLngth"`

	// Trailer field number.
	TrailerField *Number `xml:"TrlrFld,omitempty"`
}

func (p *Parameter8) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm11Code)(&value)
}

func (p *Parameter8) AddMaskGeneratorAlgorithm() *AlgorithmIdentification12 {
	p.MaskGeneratorAlgorithm = new(AlgorithmIdentification12)
	return p.MaskGeneratorAlgorithm
}

func (p *Parameter8) SetSaltLength(value string) {
	p.SaltLength = (*Number)(&value)
}

func (p *Parameter8) SetTrailerField(value string) {
	p.TrailerField = (*Number)(&value)
}
