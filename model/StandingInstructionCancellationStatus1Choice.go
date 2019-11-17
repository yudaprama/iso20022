package model

// Provides information about the status of a standing instruction cancellation request.
type StandingInstructionCancellationStatus1Choice struct {

	// Provides information about the processing status of the cancellation request.
	ProcessedStatus *CorporateActionStandingInstructionCancellationProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionStandingInstructionCancellationRejectionStatus1 `xml:"RjctdSts"`
}

func (s *StandingInstructionCancellationStatus1Choice) AddProcessedStatus() *CorporateActionStandingInstructionCancellationProcessingStatus1 {
	s.ProcessedStatus = new(CorporateActionStandingInstructionCancellationProcessingStatus1)
	return s.ProcessedStatus
}

func (s *StandingInstructionCancellationStatus1Choice) AddRejectedStatus() *CorporateActionStandingInstructionCancellationRejectionStatus1 {
	s.RejectedStatus = new(CorporateActionStandingInstructionCancellationRejectionStatus1)
	return s.RejectedStatus
}
