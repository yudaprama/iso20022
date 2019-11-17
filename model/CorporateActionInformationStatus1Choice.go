package model

// Choice between the different statuses of a corporate action information advice.
type CorporateActionInformationStatus1Choice struct {

	// Provides information about the processing status of the advice.
	ProcessedStatus *CorporateActionInformationProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionInformationRejectedStatus1 `xml:"RjctdSts"`
}

func (c *CorporateActionInformationStatus1Choice) AddProcessedStatus() *CorporateActionInformationProcessingStatus1 {
	c.ProcessedStatus = new(CorporateActionInformationProcessingStatus1)
	return c.ProcessedStatus
}

func (c *CorporateActionInformationStatus1Choice) AddRejectedStatus() *CorporateActionInformationRejectedStatus1 {
	c.RejectedStatus = new(CorporateActionInformationRejectedStatus1)
	return c.RejectedStatus
}
