package model

// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData3 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Information related to the transport key.
	Recipient []*Recipient3Choice `xml:"Rcpt"`

	// Algorithm to compute message authentication code (MAC).
	MACAlgorithm *AlgorithmIdentification10 `xml:"MACAlgo"`

	// Data to authenticate.
	EncapsulatedContent *EncapsulatedContent2 `xml:"NcpsltdCntt"`

	// Encrypted data which authenticates the data.
	MAC *Max35Binary `xml:"MAC"`
}

func (a *AuthenticatedData3) SetVersion(value string) {
	a.Version = (*Number)(&value)
}

func (a *AuthenticatedData3) AddRecipient() *Recipient3Choice {
	newValue := new(Recipient3Choice)
	a.Recipient = append(a.Recipient, newValue)
	return newValue
}

func (a *AuthenticatedData3) AddMACAlgorithm() *AlgorithmIdentification10 {
	a.MACAlgorithm = new(AlgorithmIdentification10)
	return a.MACAlgorithm
}

func (a *AuthenticatedData3) AddEncapsulatedContent() *EncapsulatedContent2 {
	a.EncapsulatedContent = new(EncapsulatedContent2)
	return a.EncapsulatedContent
}

func (a *AuthenticatedData3) SetMAC(value string) {
	a.MAC = (*Max35Binary)(&value)
}
