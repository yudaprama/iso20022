package model

// Status and reason of a transaction.
type StatusAndReason12 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus23Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus7Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus7Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus7Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason12) AddProcessingStatus() *ProcessingStatus23Choice {
	s.ProcessingStatus = new(ProcessingStatus23Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason12) AddInferredMatchingStatus() *MatchingStatus7Choice {
	s.InferredMatchingStatus = new(MatchingStatus7Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason12) AddMatchingStatus() *MatchingStatus7Choice {
	s.MatchingStatus = new(MatchingStatus7Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason12) AddSettlementStatus() *SettlementStatus7Choice {
	s.SettlementStatus = new(SettlementStatus7Choice)
	return s.SettlementStatus
}
