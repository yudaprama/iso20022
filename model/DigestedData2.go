package model

// Digest computed on the identified data.
type DigestedData2 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm.
	DigestAlgorithm []*AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Data on which the digest is computed.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Result of data-digesting process.
	Digest *Max140Text `xml:"Dgst"`
}

func (d *DigestedData2) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DigestedData2) AddDigestAlgorithm() *AlgorithmIdentification5 {
	newValue := new(AlgorithmIdentification5)
	d.DigestAlgorithm = append(d.DigestAlgorithm, newValue)
	return newValue
}

func (d *DigestedData2) AddEncapsulatedContent() *EncapsulatedContent1 {
	d.EncapsulatedContent = new(EncapsulatedContent1)
	return d.EncapsulatedContent
}

func (d *DigestedData2) SetDigest(value string) {
	d.Digest = (*Max140Text)(&value)
}
