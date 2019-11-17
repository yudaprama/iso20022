package model

// Provide processing status information of an election advice.
type CorporateActionInstructionProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus3FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionInstructionProcessingStatus1) AddStatus() *ProcessedStatus3FormatChoice {
	c.Status = new(ProcessedStatus3FormatChoice)
	return c.Status
}

func (c *CorporateActionInstructionProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
