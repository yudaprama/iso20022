package model

// Parameters of the asymmetric encryption algorithm.
type Parameter4 struct {

	// Format of data before encryption, if the format is not plaintext or implicit.
	EncryptionFormat *EncryptionFormat1Code `xml:"NcrptnFrmt,omitempty"`

	// Identification of the digest algorithm.
	DigestAlgorithm *Algorithm11Code `xml:"DgstAlgo,omitempty"`

	// Mask generator function cryptographic algorithm and parameters.
	MaskGeneratorAlgorithm *AlgorithmIdentification12 `xml:"MskGnrtrAlgo,omitempty"`
}

func (p *Parameter4) SetEncryptionFormat(value string) {
	p.EncryptionFormat = (*EncryptionFormat1Code)(&value)
}

func (p *Parameter4) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm11Code)(&value)
}

func (p *Parameter4) AddMaskGeneratorAlgorithm() *AlgorithmIdentification12 {
	p.MaskGeneratorAlgorithm = new(AlgorithmIdentification12)
	return p.MaskGeneratorAlgorithm
}
