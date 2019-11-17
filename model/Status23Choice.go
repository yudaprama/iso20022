package model

// Choice of status.
type Status23Choice struct {

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus32Choice `xml:"MtchgSts"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus32Choice `xml:"IfrrdMtchgSts"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus22Choice `xml:"SttlmSts"`

	// Provides the status of an instruction.
	InstructionProcessingStatus *InstructionProcessingStatus27Choice `xml:"InstrPrcgSts"`
}

func (s *Status23Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	s.Proprietary = new(ProprietaryStatusAndReason7)
	return s.Proprietary
}

func (s *Status23Choice) AddMatchingStatus() *MatchingStatus32Choice {
	s.MatchingStatus = new(MatchingStatus32Choice)
	return s.MatchingStatus
}

func (s *Status23Choice) AddInferredMatchingStatus() *MatchingStatus32Choice {
	s.InferredMatchingStatus = new(MatchingStatus32Choice)
	return s.InferredMatchingStatus
}

func (s *Status23Choice) AddSettlementStatus() *SettlementStatus22Choice {
	s.SettlementStatus = new(SettlementStatus22Choice)
	return s.SettlementStatus
}

func (s *Status23Choice) AddInstructionProcessingStatus() *InstructionProcessingStatus27Choice {
	s.InstructionProcessingStatus = new(InstructionProcessingStatus27Choice)
	return s.InstructionProcessingStatus
}
