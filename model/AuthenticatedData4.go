package model

// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Session key or protection key identification used by the recipient.
	Recipient []*Recipient4Choice `xml:"Rcpt"`

	// Algorithm to compute message authentication code (MAC).
	MACAlgorithm *AlgorithmIdentification15 `xml:"MACAlgo"`

	// Data to authenticate.
	EncapsulatedContent *EncapsulatedContent3 `xml:"NcpsltdCntt"`

	// Message authentication code value.
	MAC *Max140Binary `xml:"MAC"`
}

func (a *AuthenticatedData4) SetVersion(value string) {
	a.Version = (*Number)(&value)
}

func (a *AuthenticatedData4) AddRecipient() *Recipient4Choice {
	newValue := new(Recipient4Choice)
	a.Recipient = append(a.Recipient, newValue)
	return newValue
}

func (a *AuthenticatedData4) AddMACAlgorithm() *AlgorithmIdentification15 {
	a.MACAlgorithm = new(AlgorithmIdentification15)
	return a.MACAlgorithm
}

func (a *AuthenticatedData4) AddEncapsulatedContent() *EncapsulatedContent3 {
	a.EncapsulatedContent = new(EncapsulatedContent3)
	return a.EncapsulatedContent
}

func (a *AuthenticatedData4) SetMAC(value string) {
	a.MAC = (*Max140Binary)(&value)
}
