package model

// Provide processing status information of a standing instruction cancellation request.
type CorporateActionStandingInstructionCancellationProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus4FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionStandingInstructionCancellationProcessingStatus1) AddStatus() *ProcessedStatus4FormatChoice {
	c.Status = new(ProcessedStatus4FormatChoice)
	return c.Status
}

func (c *CorporateActionStandingInstructionCancellationProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
