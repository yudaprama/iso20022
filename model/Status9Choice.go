package model

// Choice of status.
type Status9Choice struct {

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus7Choice `xml:"MtchgSts"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus7Choice `xml:"IfrrdMtchgSts"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus7Choice `xml:"SttlmSts"`

	// Provides the status of an instruction.
	InstructionProcessingStatus *InstructionProcessingStatus10Choice `xml:"InstrPrcgSts"`
}

func (s *Status9Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	s.Proprietary = new(ProprietaryStatusAndReason1)
	return s.Proprietary
}

func (s *Status9Choice) AddMatchingStatus() *MatchingStatus7Choice {
	s.MatchingStatus = new(MatchingStatus7Choice)
	return s.MatchingStatus
}

func (s *Status9Choice) AddInferredMatchingStatus() *MatchingStatus7Choice {
	s.InferredMatchingStatus = new(MatchingStatus7Choice)
	return s.InferredMatchingStatus
}

func (s *Status9Choice) AddSettlementStatus() *SettlementStatus7Choice {
	s.SettlementStatus = new(SettlementStatus7Choice)
	return s.SettlementStatus
}

func (s *Status9Choice) AddInstructionProcessingStatus() *InstructionProcessingStatus10Choice {
	s.InstructionProcessingStatus = new(InstructionProcessingStatus10Choice)
	return s.InstructionProcessingStatus
}
