package model

// Provide processing status information of an election advice.
type CorporateActionCancellationProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus5FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionCancellationProcessingStatus1) AddStatus() *ProcessedStatus5FormatChoice {
	c.Status = new(ProcessedStatus5FormatChoice)
	return c.Status
}

func (c *CorporateActionCancellationProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
