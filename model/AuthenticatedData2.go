package model

// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData2 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Information related to the transport key.
	Recipient []*Recipient2Choice `xml:"Rcpt"`

	// Algorithm to compute message authentication code (MAC).
	MACAlgorithm *AlgorithmIdentification3 `xml:"MACAlgo"`

	// Data to authenticate.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Encrypted data which authenticates the data.
	MAC *Max35Binary `xml:"MAC"`
}

func (a *AuthenticatedData2) SetVersion(value string) {
	a.Version = (*Number)(&value)
}

func (a *AuthenticatedData2) AddRecipient() *Recipient2Choice {
	newValue := new(Recipient2Choice)
	a.Recipient = append(a.Recipient, newValue)
	return newValue
}

func (a *AuthenticatedData2) AddMACAlgorithm() *AlgorithmIdentification3 {
	a.MACAlgorithm = new(AlgorithmIdentification3)
	return a.MACAlgorithm
}

func (a *AuthenticatedData2) AddEncapsulatedContent() *EncapsulatedContent1 {
	a.EncapsulatedContent = new(EncapsulatedContent1)
	return a.EncapsulatedContent
}

func (a *AuthenticatedData2) SetMAC(value string) {
	a.MAC = (*Max35Binary)(&value)
}
