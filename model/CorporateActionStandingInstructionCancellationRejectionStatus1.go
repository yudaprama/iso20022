package model

// Provides reason of the rejection of a standing instruction cancellation request.
type CorporateActionStandingInstructionCancellationRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason10FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionStandingInstructionCancellationRejectionStatus1) AddReason() *RejectionReason10FormatChoice {
	newValue := new(RejectionReason10FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionStandingInstructionCancellationRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
