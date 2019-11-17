package model

// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData1 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Information related to the transport key.
	Recipient []*Recipient1Choice `xml:"Rcpt"`

	// Algorithm to compute message authentication code (MAC).
	MACAlgorithm *AlgorithmIdentification1 `xml:"MACAlgo"`

	// Data to authenticate.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Encrypted data which authenticates the data.
	MAC *Max35Binary `xml:"MAC"`
}

func (a *AuthenticatedData1) SetVersion(value string) {
	a.Version = (*Number)(&value)
}

func (a *AuthenticatedData1) AddRecipient() *Recipient1Choice {
	newValue := new(Recipient1Choice)
	a.Recipient = append(a.Recipient, newValue)
	return newValue
}

func (a *AuthenticatedData1) AddMACAlgorithm() *AlgorithmIdentification1 {
	a.MACAlgorithm = new(AlgorithmIdentification1)
	return a.MACAlgorithm
}

func (a *AuthenticatedData1) AddEncapsulatedContent() *EncapsulatedContent1 {
	a.EncapsulatedContent = new(EncapsulatedContent1)
	return a.EncapsulatedContent
}

func (a *AuthenticatedData1) SetMAC(value string) {
	a.MAC = (*Max35Binary)(&value)
}
