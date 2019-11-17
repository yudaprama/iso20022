package model

// Status and reason of a transaction.
type StatusAndReason19 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus23Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus19Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus19Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus7Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason19) AddProcessingStatus() *ProcessingStatus23Choice {
	s.ProcessingStatus = new(ProcessingStatus23Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason19) AddInferredMatchingStatus() *MatchingStatus19Choice {
	s.InferredMatchingStatus = new(MatchingStatus19Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason19) AddMatchingStatus() *MatchingStatus19Choice {
	s.MatchingStatus = new(MatchingStatus19Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason19) AddSettlementStatus() *SettlementStatus7Choice {
	s.SettlementStatus = new(SettlementStatus7Choice)
	return s.SettlementStatus
}
