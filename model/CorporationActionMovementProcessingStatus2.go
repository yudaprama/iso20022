package model

// Provides processing status of the movement cancellation request.
type CorporationActionMovementProcessingStatus2 struct {

	// The processing status.
	Status *ProcessedStatus2FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporationActionMovementProcessingStatus2) AddStatus() *ProcessedStatus2FormatChoice {
	c.Status = new(ProcessedStatus2FormatChoice)
	return c.Status
}

func (c *CorporationActionMovementProcessingStatus2) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
