package model

// Provides reason of the rejection of a deactivation instruction.
type CorporateActionDeactivationInstructionRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason12FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionDeactivationInstructionRejectionStatus1) AddReason() *RejectionReason12FormatChoice {
	newValue := new(RejectionReason12FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionDeactivationInstructionRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
