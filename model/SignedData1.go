package model

// Digital signature of data, with an asymmetric key.
type SignedData1 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm []*AlgorithmIdentification1 `xml:"DgstAlgo"`

	// Data to sign.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Chain of X.509 certificates.
	Certificate []*Max3000Binary `xml:"Cert,omitempty"`

	// Entity who has signed the data.
	Signer []*Signer1 `xml:"Sgnr"`
}

func (s *SignedData1) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *SignedData1) AddDigestAlgorithm() *AlgorithmIdentification1 {
	newValue := new(AlgorithmIdentification1)
	s.DigestAlgorithm = append(s.DigestAlgorithm, newValue)
	return newValue
}

func (s *SignedData1) AddEncapsulatedContent() *EncapsulatedContent1 {
	s.EncapsulatedContent = new(EncapsulatedContent1)
	return s.EncapsulatedContent
}

func (s *SignedData1) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max3000Binary)(&value))
}

func (s *SignedData1) AddSigner() *Signer1 {
	newValue := new(Signer1)
	s.Signer = append(s.Signer, newValue)
	return newValue
}
