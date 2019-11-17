package model

// Choice of status.
type Status18Choice struct {

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus24Choice `xml:"MtchgSts"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus24Choice `xml:"IfrrdMtchgSts"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus17Choice `xml:"SttlmSts"`

	// Provides the status of an instruction.
	InstructionProcessingStatus *InstructionProcessingStatus22Choice `xml:"InstrPrcgSts"`
}

func (s *Status18Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	s.Proprietary = new(ProprietaryStatusAndReason6)
	return s.Proprietary
}

func (s *Status18Choice) AddMatchingStatus() *MatchingStatus24Choice {
	s.MatchingStatus = new(MatchingStatus24Choice)
	return s.MatchingStatus
}

func (s *Status18Choice) AddInferredMatchingStatus() *MatchingStatus24Choice {
	s.InferredMatchingStatus = new(MatchingStatus24Choice)
	return s.InferredMatchingStatus
}

func (s *Status18Choice) AddSettlementStatus() *SettlementStatus17Choice {
	s.SettlementStatus = new(SettlementStatus17Choice)
	return s.SettlementStatus
}

func (s *Status18Choice) AddInstructionProcessingStatus() *InstructionProcessingStatus22Choice {
	s.InstructionProcessingStatus = new(InstructionProcessingStatus22Choice)
	return s.InstructionProcessingStatus
}
