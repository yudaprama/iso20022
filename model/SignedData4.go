package model

// Digital signatures of data from one or several signers.
type SignedData4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of digest algorithm applied before signature.
	DigestAlgorithm []*AlgorithmIdentification16 `xml:"DgstAlgo"`

	// Data to sign.
	EncapsulatedContent *EncapsulatedContent3 `xml:"NcpsltdCntt"`

	// Chain of X.509 certificates.
	Certificate []*Max5000Binary `xml:"Cert,omitempty"`

	// Digital signature and identification of a signer.
	Signer []*Signer3 `xml:"Sgnr"`
}

func (s *SignedData4) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *SignedData4) AddDigestAlgorithm() *AlgorithmIdentification16 {
	newValue := new(AlgorithmIdentification16)
	s.DigestAlgorithm = append(s.DigestAlgorithm, newValue)
	return newValue
}

func (s *SignedData4) AddEncapsulatedContent() *EncapsulatedContent3 {
	s.EncapsulatedContent = new(EncapsulatedContent3)
	return s.EncapsulatedContent
}

func (s *SignedData4) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max5000Binary)(&value))
}

func (s *SignedData4) AddSigner() *Signer3 {
	newValue := new(Signer3)
	s.Signer = append(s.Signer, newValue)
	return newValue
}
