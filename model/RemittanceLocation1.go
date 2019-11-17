package model

// Remittance information that provides all remittance address elements, that enables the matching, i.e.  reconciliation, of a payment with the items that the transaction in intended to settle.
type RemittanceLocation1 struct {

	// Unique and unambiguous identification of the remittance information, e.g. a remittance advice, which is sent separately from the payment instruction.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Specifies the method used to deliver the remittance advice information.
	RemittanceLocationMethod *RemittanceLocationMethod1Code `xml:"RmtLctnMtd,omitempty"`

	// Electronic address to which an agent is to send the remittance information.
	RemittanceLocationElectronicAddress *Max256Text `xml:"RmtLctnElctrncAdr,omitempty"`

	// Postal address to which an agent is to send the remittance information.
	RemittanceLocationPostalAddress *NameAndAddress3 `xml:"RmtLctnPstlAdr,omitempty"`
}

func (r *RemittanceLocation1) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceLocation1) SetRemittanceLocationMethod(value string) {
	r.RemittanceLocationMethod = (*RemittanceLocationMethod1Code)(&value)
}

func (r *RemittanceLocation1) SetRemittanceLocationElectronicAddress(value string) {
	r.RemittanceLocationElectronicAddress = (*Max256Text)(&value)
}

func (r *RemittanceLocation1) AddRemittanceLocationPostalAddress() *NameAndAddress3 {
	r.RemittanceLocationPostalAddress = new(NameAndAddress3)
	return r.RemittanceLocationPostalAddress
}
