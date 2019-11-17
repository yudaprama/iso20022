package model

// Provide processing status information of an election amendment request.
type CorporateActionAmendmentProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus5FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionAmendmentProcessingStatus1) AddStatus() *ProcessedStatus5FormatChoice {
	c.Status = new(ProcessedStatus5FormatChoice)
	return c.Status
}

func (c *CorporateActionAmendmentProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
