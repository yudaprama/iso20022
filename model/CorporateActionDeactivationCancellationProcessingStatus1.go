package model

// Provide the processing status information for a deactivation cancellation request.
type CorporateActionDeactivationCancellationProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus2FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionDeactivationCancellationProcessingStatus1) AddStatus() *ProcessedStatus2FormatChoice {
	c.Status = new(ProcessedStatus2FormatChoice)
	return c.Status
}

func (c *CorporateActionDeactivationCancellationProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
