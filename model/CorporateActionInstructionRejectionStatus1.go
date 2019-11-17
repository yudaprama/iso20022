package model

// Provides reason of the rejection of an election advice.
type CorporateActionInstructionRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason18FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionInstructionRejectionStatus1) AddReason() *RejectionReason18FormatChoice {
	newValue := new(RejectionReason18FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionInstructionRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
