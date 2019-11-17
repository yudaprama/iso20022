package model

// Choice between the different statuses of an election advice.
type ElectionAdviceStatus1Choice struct {

	// Provides information about the processing status of advice.
	ProcessedStatus *CorporateActionInstructionProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionInstructionRejectionStatus1 `xml:"RjctdSts"`
}

func (e *ElectionAdviceStatus1Choice) AddProcessedStatus() *CorporateActionInstructionProcessingStatus1 {
	e.ProcessedStatus = new(CorporateActionInstructionProcessingStatus1)
	return e.ProcessedStatus
}

func (e *ElectionAdviceStatus1Choice) AddRejectedStatus() *CorporateActionInstructionRejectionStatus1 {
	e.RejectedStatus = new(CorporateActionInstructionRejectionStatus1)
	return e.RejectedStatus
}
