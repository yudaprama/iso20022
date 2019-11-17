package model

// Set of elements used to provide information on the remittance advice.
type RemittanceLocation4 struct {

	// Unique identification, as assigned by the initiating party, to unambiguously identify the remittance information sent separately from the payment instruction, such as a remittance advice.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Set of elements used to provide information on the location and/or delivery of the remittance information.
	RemittanceLocationDetails []*RemittanceLocationDetails1 `xml:"RmtLctnDtls,omitempty"`
}

func (r *RemittanceLocation4) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceLocation4) AddRemittanceLocationDetails() *RemittanceLocationDetails1 {
	newValue := new(RemittanceLocationDetails1)
	r.RemittanceLocationDetails = append(r.RemittanceLocationDetails, newValue)
	return newValue
}
