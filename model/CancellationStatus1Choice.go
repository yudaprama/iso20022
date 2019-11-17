package model

// Status applying to the instruction cancellation request received.
type CancellationStatus1Choice struct {

	// Status advising on the processing of the cancellation request.
	ProcessingStatus *CancellationProcessingStatus1 `xml:"PrcgSts"`

	// Status advising on the rejection of the cancellation request.
	RejectionStatus *CancellationRejectionStatus1 `xml:"RjctnSts"`
}

func (c *CancellationStatus1Choice) AddProcessingStatus() *CancellationProcessingStatus1 {
	c.ProcessingStatus = new(CancellationProcessingStatus1)
	return c.ProcessingStatus
}

func (c *CancellationStatus1Choice) AddRejectionStatus() *CancellationRejectionStatus1 {
	c.RejectionStatus = new(CancellationRejectionStatus1)
	return c.RejectionStatus
}
