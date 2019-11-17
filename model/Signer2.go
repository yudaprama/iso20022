package model

// Entity who has signed the data and its digital signature.
type Signer2 struct {

	// Version of the Cryptographic Message Syntax (CMS) data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the entity who has signed the data.
	SignerIdentification *CertificateIdentifier1 `xml:"SgnrId"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm *AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Cryptographic digital signature algorithm.
	SignatureAlgorithm *AlgorithmIdentification4 `xml:"SgntrAlgo"`

	// Digital signature.
	Signature *Max500Binary `xml:"Sgntr"`
}

func (s *Signer2) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *Signer2) AddSignerIdentification() *CertificateIdentifier1 {
	s.SignerIdentification = new(CertificateIdentifier1)
	return s.SignerIdentification
}

func (s *Signer2) AddDigestAlgorithm() *AlgorithmIdentification5 {
	s.DigestAlgorithm = new(AlgorithmIdentification5)
	return s.DigestAlgorithm
}

func (s *Signer2) AddSignatureAlgorithm() *AlgorithmIdentification4 {
	s.SignatureAlgorithm = new(AlgorithmIdentification4)
	return s.SignatureAlgorithm
}

func (s *Signer2) SetSignature(value string) {
	s.Signature = (*Max500Binary)(&value)
}
