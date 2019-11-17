package model

// Digest computed on the identified data.
type DigestedData3 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm.
	DigestAlgorithm []*AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Data on which the digest is computed.
	EncapsulatedContent *EncapsulatedContent2 `xml:"NcpsltdCntt"`

	// Result of data-digesting process.
	Digest *Max140Text `xml:"Dgst"`
}

func (d *DigestedData3) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DigestedData3) AddDigestAlgorithm() *AlgorithmIdentification5 {
	newValue := new(AlgorithmIdentification5)
	d.DigestAlgorithm = append(d.DigestAlgorithm, newValue)
	return newValue
}

func (d *DigestedData3) AddEncapsulatedContent() *EncapsulatedContent2 {
	d.EncapsulatedContent = new(EncapsulatedContent2)
	return d.EncapsulatedContent
}

func (d *DigestedData3) SetDigest(value string) {
	d.Digest = (*Max140Text)(&value)
}
