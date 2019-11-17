package model

// Choice of status.
type Status15Choice struct {

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus19Choice `xml:"MtchgSts"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus19Choice `xml:"IfrrdMtchgSts"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus7Choice `xml:"SttlmSts"`

	// Provides the status of an instruction.
	InstructionProcessingStatus *InstructionProcessingStatus14Choice `xml:"InstrPrcgSts"`
}

func (s *Status15Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	s.Proprietary = new(ProprietaryStatusAndReason1)
	return s.Proprietary
}

func (s *Status15Choice) AddMatchingStatus() *MatchingStatus19Choice {
	s.MatchingStatus = new(MatchingStatus19Choice)
	return s.MatchingStatus
}

func (s *Status15Choice) AddInferredMatchingStatus() *MatchingStatus19Choice {
	s.InferredMatchingStatus = new(MatchingStatus19Choice)
	return s.InferredMatchingStatus
}

func (s *Status15Choice) AddSettlementStatus() *SettlementStatus7Choice {
	s.SettlementStatus = new(SettlementStatus7Choice)
	return s.SettlementStatus
}

func (s *Status15Choice) AddInstructionProcessingStatus() *InstructionProcessingStatus14Choice {
	s.InstructionProcessingStatus = new(InstructionProcessingStatus14Choice)
	return s.InstructionProcessingStatus
}
