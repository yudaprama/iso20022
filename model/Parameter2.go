package model

// Parameters of the RSAES-OAEP encryption algorithm (RSA Encryption Scheme: Optimal Asymmetric Encryption Padding).
type Parameter2 struct {

	// Digest algorithm used in the RSAES-OAEP encryption algorithm (RSA Encryption Scheme: Optimal Asymmetric Encryption Padding).
	DigestAlgorithm *Algorithm5Code `xml:"DgstAlgo,omitempty"`

	// Mask generator function cryptographic algorithm and parameters.
	MaskGeneratorAlgorithm *AlgorithmIdentification8 `xml:"MskGnrtrAlgo,omitempty"`
}

func (p *Parameter2) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm5Code)(&value)
}

func (p *Parameter2) AddMaskGeneratorAlgorithm() *AlgorithmIdentification8 {
	p.MaskGeneratorAlgorithm = new(AlgorithmIdentification8)
	return p.MaskGeneratorAlgorithm
}
