package model

// Digest computed on the identified data.
type DigestedData4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the digest algorithm.
	DigestAlgorithm *AlgorithmIdentification16 `xml:"DgstAlgo"`

	// Data on which the digest is computed.
	EncapsulatedContent *EncapsulatedContent3 `xml:"NcpsltdCntt"`

	// Result of data-digesting process.
	Digest *Max140Binary `xml:"Dgst"`
}

func (d *DigestedData4) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DigestedData4) AddDigestAlgorithm() *AlgorithmIdentification16 {
	d.DigestAlgorithm = new(AlgorithmIdentification16)
	return d.DigestAlgorithm
}

func (d *DigestedData4) AddEncapsulatedContent() *EncapsulatedContent3 {
	d.EncapsulatedContent = new(EncapsulatedContent3)
	return d.EncapsulatedContent
}

func (d *DigestedData4) SetDigest(value string) {
	d.Digest = (*Max140Binary)(&value)
}
