package model

// Status and reason of a transaction.
type StatusAndReason29 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus62Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus32Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus32Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus22Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason29) AddProcessingStatus() *ProcessingStatus62Choice {
	s.ProcessingStatus = new(ProcessingStatus62Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason29) AddInferredMatchingStatus() *MatchingStatus32Choice {
	s.InferredMatchingStatus = new(MatchingStatus32Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason29) AddMatchingStatus() *MatchingStatus32Choice {
	s.MatchingStatus = new(MatchingStatus32Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason29) AddSettlementStatus() *SettlementStatus22Choice {
	s.SettlementStatus = new(SettlementStatus22Choice)
	return s.SettlementStatus
}
