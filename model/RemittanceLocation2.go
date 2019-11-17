package model

// Set of elements used to provide information on the remittance advice.
type RemittanceLocation2 struct {

	// Unique identification, as assigned by the initiating party, to unambiguously identify the remittance information sent separately from the payment instruction, such as a remittance advice.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Method used to deliver the remittance advice information.
	RemittanceLocationMethod *RemittanceLocationMethod2Code `xml:"RmtLctnMtd,omitempty"`

	// Electronic address to which an agent is to send the remittance information.
	RemittanceLocationElectronicAddress *Max2048Text `xml:"RmtLctnElctrncAdr,omitempty"`

	// Postal address to which an agent is to send the remittance information.
	RemittanceLocationPostalAddress *NameAndAddress10 `xml:"RmtLctnPstlAdr,omitempty"`
}

func (r *RemittanceLocation2) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceLocation2) SetRemittanceLocationMethod(value string) {
	r.RemittanceLocationMethod = (*RemittanceLocationMethod2Code)(&value)
}

func (r *RemittanceLocation2) SetRemittanceLocationElectronicAddress(value string) {
	r.RemittanceLocationElectronicAddress = (*Max2048Text)(&value)
}

func (r *RemittanceLocation2) AddRemittanceLocationPostalAddress() *NameAndAddress10 {
	r.RemittanceLocationPostalAddress = new(NameAndAddress10)
	return r.RemittanceLocationPostalAddress
}
