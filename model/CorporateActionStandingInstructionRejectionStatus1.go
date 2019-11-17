package model

// Provides reason of the rejection of a standing instruction request.
type CorporateActionStandingInstructionRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason20FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionStandingInstructionRejectionStatus1) AddReason() *RejectionReason20FormatChoice {
	newValue := new(RejectionReason20FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionStandingInstructionRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
