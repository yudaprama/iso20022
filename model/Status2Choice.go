package model

// Choice of status.
type Status2Choice struct {

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus2Choice `xml:"MtchgSts"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus2Choice `xml:"IfrrdMtchgSts"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus2Choice `xml:"SttlmSts"`

	// Provides the status of an instruction.
	InstructionProcessingStatus *InstructionProcessingStatus3Choice `xml:"InstrPrcgSts"`
}

func (s *Status2Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	s.Proprietary = new(ProprietaryStatusAndReason1)
	return s.Proprietary
}

func (s *Status2Choice) AddMatchingStatus() *MatchingStatus2Choice {
	s.MatchingStatus = new(MatchingStatus2Choice)
	return s.MatchingStatus
}

func (s *Status2Choice) AddInferredMatchingStatus() *MatchingStatus2Choice {
	s.InferredMatchingStatus = new(MatchingStatus2Choice)
	return s.InferredMatchingStatus
}

func (s *Status2Choice) AddSettlementStatus() *SettlementStatus2Choice {
	s.SettlementStatus = new(SettlementStatus2Choice)
	return s.SettlementStatus
}

func (s *Status2Choice) AddInstructionProcessingStatus() *InstructionProcessingStatus3Choice {
	s.InstructionProcessingStatus = new(InstructionProcessingStatus3Choice)
	return s.InstructionProcessingStatus
}
