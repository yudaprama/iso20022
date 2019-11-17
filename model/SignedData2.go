package model

// Digital signature of data, with an asymmetric key.
type SignedData2 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm []*AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Data to sign.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Chain of X.509 certificates.
	Certificate []*Max3000Binary `xml:"Cert,omitempty"`

	// Entity who has signed the data.
	Signer []*Signer2 `xml:"Sgnr"`
}

func (s *SignedData2) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *SignedData2) AddDigestAlgorithm() *AlgorithmIdentification5 {
	newValue := new(AlgorithmIdentification5)
	s.DigestAlgorithm = append(s.DigestAlgorithm, newValue)
	return newValue
}

func (s *SignedData2) AddEncapsulatedContent() *EncapsulatedContent1 {
	s.EncapsulatedContent = new(EncapsulatedContent1)
	return s.EncapsulatedContent
}

func (s *SignedData2) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max3000Binary)(&value))
}

func (s *SignedData2) AddSigner() *Signer2 {
	newValue := new(Signer2)
	s.Signer = append(s.Signer, newValue)
	return newValue
}
