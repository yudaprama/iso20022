package model

// Provide the processing status information of information advice.
type CorporateActionInformationProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus5FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionInformationProcessingStatus1) AddStatus() *ProcessedStatus5FormatChoice {
	c.Status = new(ProcessedStatus5FormatChoice)
	return c.Status
}

func (c *CorporateActionInformationProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
