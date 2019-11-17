package model

// Choice between the different statuses of an election amendment request.
type ElectionAmendmentStatus1Choice struct {

	// Provides information about the processing status of the request.
	ProcessedStatus *CorporateActionAmendmentProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionAmendmentRejectionStatus1 `xml:"RjctdSts"`
}

func (e *ElectionAmendmentStatus1Choice) AddProcessedStatus() *CorporateActionAmendmentProcessingStatus1 {
	e.ProcessedStatus = new(CorporateActionAmendmentProcessingStatus1)
	return e.ProcessedStatus
}

func (e *ElectionAmendmentStatus1Choice) AddRejectedStatus() *CorporateActionAmendmentRejectionStatus1 {
	e.RejectedStatus = new(CorporateActionAmendmentRejectionStatus1)
	return e.RejectedStatus
}
