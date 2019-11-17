package model

// Status advising on the processing of the cancellation request.
type CancellationProcessingStatus1 struct {

	// Status of the request for cancellation.
	Status *CancellationStatus3Code `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationProcessingStatus1) SetStatus(value string) {
	c.Status = (*CancellationStatus3Code)(&value)
}

func (c *CancellationProcessingStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
