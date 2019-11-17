package model

// Choice between the different statuses of a movement.
type CorporateActionMovementStatus1Choice struct {

	// Provides information about the processing status of the movement.
	ProcessedStatus *CorporateActionMovementProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the settlement failure.
	FailedStatus *CorporateActionMovementFailedStatus1 `xml:"FaildSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionMovementRejectionStatus1 `xml:"RjctdSts"`
}

func (c *CorporateActionMovementStatus1Choice) AddProcessedStatus() *CorporateActionMovementProcessingStatus1 {
	c.ProcessedStatus = new(CorporateActionMovementProcessingStatus1)
	return c.ProcessedStatus
}

func (c *CorporateActionMovementStatus1Choice) AddFailedStatus() *CorporateActionMovementFailedStatus1 {
	c.FailedStatus = new(CorporateActionMovementFailedStatus1)
	return c.FailedStatus
}

func (c *CorporateActionMovementStatus1Choice) AddRejectedStatus() *CorporateActionMovementRejectionStatus1 {
	c.RejectedStatus = new(CorporateActionMovementRejectionStatus1)
	return c.RejectedStatus
}
