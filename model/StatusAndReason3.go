package model

// Status and reason of a transaction.
type StatusAndReason3 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus6Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus2Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus2Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus2Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason3) AddProcessingStatus() *ProcessingStatus6Choice {
	s.ProcessingStatus = new(ProcessingStatus6Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason3) AddInferredMatchingStatus() *MatchingStatus2Choice {
	s.InferredMatchingStatus = new(MatchingStatus2Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason3) AddMatchingStatus() *MatchingStatus2Choice {
	s.MatchingStatus = new(MatchingStatus2Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason3) AddSettlementStatus() *SettlementStatus2Choice {
	s.SettlementStatus = new(SettlementStatus2Choice)
	return s.SettlementStatus
}
