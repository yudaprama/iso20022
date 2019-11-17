package model

// Choice of format for the settlement status.
type SettlementStatus17Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus37Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus10Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (s *SettlementStatus17Choice) AddPending() *PendingStatus37Choice {
	s.Pending = new(PendingStatus37Choice)
	return s.Pending
}

func (s *SettlementStatus17Choice) AddFailing() *FailingStatus10Choice {
	s.Failing = new(FailingStatus10Choice)
	return s.Failing
}

func (s *SettlementStatus17Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	s.Proprietary = new(ProprietaryStatusAndReason6)
	return s.Proprietary
}
