package model

// Choice between various statuses.
type StandingInstructionStatus1Choice struct {

	// Provides information about the processing status of the request.
	ProcessedStatus *CorporateActionStandingInstructionProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionStandingInstructionRejectionStatus1 `xml:"RjctdSts"`
}

func (s *StandingInstructionStatus1Choice) AddProcessedStatus() *CorporateActionStandingInstructionProcessingStatus1 {
	s.ProcessedStatus = new(CorporateActionStandingInstructionProcessingStatus1)
	return s.ProcessedStatus
}

func (s *StandingInstructionStatus1Choice) AddRejectedStatus() *CorporateActionStandingInstructionRejectionStatus1 {
	s.RejectedStatus = new(CorporateActionStandingInstructionRejectionStatus1)
	return s.RejectedStatus
}
