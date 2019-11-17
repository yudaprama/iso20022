package model

// Status applying to the instruction cancellation request received.
type CancellationStatus2Choice struct {

	// Status advising on the processing of the cancellation request.
	ProcessingStatus *CancellationProcessingStatus1 `xml:"PrcgSts"`

	// Status advising on the rejection of the cancellation request.
	RejectionStatus *AdditionalStatus2 `xml:"RjctnSts"`
}

func (c *CancellationStatus2Choice) AddProcessingStatus() *CancellationProcessingStatus1 {
	c.ProcessingStatus = new(CancellationProcessingStatus1)
	return c.ProcessingStatus
}

func (c *CancellationStatus2Choice) AddRejectionStatus() *AdditionalStatus2 {
	c.RejectionStatus = new(AdditionalStatus2)
	return c.RejectionStatus
}
