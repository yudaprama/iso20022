package model

// Provides status of the movement.
type CorporateMovementStatus2 struct {

	// Provides information about the processing status of the cancellation request.
	ProcessedStatus *CorporationActionMovementProcessingStatus2 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionMovementRejectionStatus2 `xml:"RjctdSts"`
}

func (c *CorporateMovementStatus2) AddProcessedStatus() *CorporationActionMovementProcessingStatus2 {
	c.ProcessedStatus = new(CorporationActionMovementProcessingStatus2)
	return c.ProcessedStatus
}

func (c *CorporateMovementStatus2) AddRejectedStatus() *CorporateActionMovementRejectionStatus2 {
	c.RejectedStatus = new(CorporateActionMovementRejectionStatus2)
	return c.RejectedStatus
}
