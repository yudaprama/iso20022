package model

// Provide processing status information of a standing instruction request.
type CorporateActionStandingInstructionProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus3FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionStandingInstructionProcessingStatus1) AddStatus() *ProcessedStatus3FormatChoice {
	c.Status = new(ProcessedStatus3FormatChoice)
	return c.Status
}

func (c *CorporateActionStandingInstructionProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
