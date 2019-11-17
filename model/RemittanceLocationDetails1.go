package model

// Provides information on the remittance advice.
type RemittanceLocationDetails1 struct {

	// Method used to deliver the remittance advice information.
	Method *RemittanceLocationMethod2Code `xml:"Mtd"`

	// Electronic address to which an agent is to send the remittance information.
	ElectronicAddress *Max2048Text `xml:"ElctrncAdr,omitempty"`

	// Postal address to which an agent is to send the remittance information.
	PostalAddress *NameAndAddress10 `xml:"PstlAdr,omitempty"`
}

func (r *RemittanceLocationDetails1) SetMethod(value string) {
	r.Method = (*RemittanceLocationMethod2Code)(&value)
}

func (r *RemittanceLocationDetails1) SetElectronicAddress(value string) {
	r.ElectronicAddress = (*Max2048Text)(&value)
}

func (r *RemittanceLocationDetails1) AddPostalAddress() *NameAndAddress10 {
	r.PostalAddress = new(NameAndAddress10)
	return r.PostalAddress
}
