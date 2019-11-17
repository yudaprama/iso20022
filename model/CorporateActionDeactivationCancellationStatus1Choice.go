package model

// Choice between the different statuses of a deactivation cancellation request.
type CorporateActionDeactivationCancellationStatus1Choice struct {

	// Provides information about the processing status of the cancellation request.
	ProcessedStatus *CorporateActionDeactivationCancellationProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionDeactivationCancellationRejectionStatus1 `xml:"RjctdSts"`
}

func (c *CorporateActionDeactivationCancellationStatus1Choice) AddProcessedStatus() *CorporateActionDeactivationCancellationProcessingStatus1 {
	c.ProcessedStatus = new(CorporateActionDeactivationCancellationProcessingStatus1)
	return c.ProcessedStatus
}

func (c *CorporateActionDeactivationCancellationStatus1Choice) AddRejectedStatus() *CorporateActionDeactivationCancellationRejectionStatus1 {
	c.RejectedStatus = new(CorporateActionDeactivationCancellationRejectionStatus1)
	return c.RejectedStatus
}
