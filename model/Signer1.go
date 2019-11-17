package model

// Entity who has signed the data and its digital signature.
type Signer1 struct {

	// Version of the Cryptographic Message Syntax (CMS) data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the entity who has signed the data.
	SignerIdentification *CertificateIdentifier1 `xml:"SgnrId"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm *AlgorithmIdentification1 `xml:"DgstAlgo"`

	// Cryptographic digital signature algorithm.
	SignatureAlgorithm *AlgorithmIdentification1 `xml:"SgntrAlgo"`

	// Digital signature.
	Signature *Max500Binary `xml:"Sgntr"`
}

func (s *Signer1) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *Signer1) AddSignerIdentification() *CertificateIdentifier1 {
	s.SignerIdentification = new(CertificateIdentifier1)
	return s.SignerIdentification
}

func (s *Signer1) AddDigestAlgorithm() *AlgorithmIdentification1 {
	s.DigestAlgorithm = new(AlgorithmIdentification1)
	return s.DigestAlgorithm
}

func (s *Signer1) AddSignatureAlgorithm() *AlgorithmIdentification1 {
	s.SignatureAlgorithm = new(AlgorithmIdentification1)
	return s.SignatureAlgorithm
}

func (s *Signer1) SetSignature(value string) {
	s.Signature = (*Max500Binary)(&value)
}
