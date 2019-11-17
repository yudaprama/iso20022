package model

// Entity who has signed the data and its digital signature.
type Signer3 struct {

	// Version of the Cryptographic Message Syntax (CMS) data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the entity who has signed the data.
	SignerIdentification *Recipient5Choice `xml:"SgnrId,omitempty"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm *AlgorithmIdentification16 `xml:"DgstAlgo"`

	// Cryptographic digital signature algorithm.
	SignatureAlgorithm *AlgorithmIdentification17 `xml:"SgntrAlgo"`

	// Digital signature.
	Signature *Max3000Binary `xml:"Sgntr"`
}

func (s *Signer3) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *Signer3) AddSignerIdentification() *Recipient5Choice {
	s.SignerIdentification = new(Recipient5Choice)
	return s.SignerIdentification
}

func (s *Signer3) AddDigestAlgorithm() *AlgorithmIdentification16 {
	s.DigestAlgorithm = new(AlgorithmIdentification16)
	return s.DigestAlgorithm
}

func (s *Signer3) AddSignatureAlgorithm() *AlgorithmIdentification17 {
	s.SignatureAlgorithm = new(AlgorithmIdentification17)
	return s.SignatureAlgorithm
}

func (s *Signer3) SetSignature(value string) {
	s.Signature = (*Max3000Binary)(&value)
}
