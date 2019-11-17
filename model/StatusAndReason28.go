package model

// Status and reason of a transaction.
type StatusAndReason28 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus52Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus24Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus24Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus17Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason28) AddProcessingStatus() *ProcessingStatus52Choice {
	s.ProcessingStatus = new(ProcessingStatus52Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason28) AddInferredMatchingStatus() *MatchingStatus24Choice {
	s.InferredMatchingStatus = new(MatchingStatus24Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason28) AddMatchingStatus() *MatchingStatus24Choice {
	s.MatchingStatus = new(MatchingStatus24Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason28) AddSettlementStatus() *SettlementStatus17Choice {
	s.SettlementStatus = new(SettlementStatus17Choice)
	return s.SettlementStatus
}
