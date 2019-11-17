package model

// Provide the processing status information of a deactivation instruction.
type CorporateActionDeactivationInstructionProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus6FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionDeactivationInstructionProcessingStatus1) AddStatus() *ProcessedStatus6FormatChoice {
	c.Status = new(ProcessedStatus6FormatChoice)
	return c.Status
}

func (c *CorporateActionDeactivationInstructionProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
