package model

// Choice of format for the settlement status.
type SettlementStatus16Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus36Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus9Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (s *SettlementStatus16Choice) AddPending() *PendingStatus36Choice {
	s.Pending = new(PendingStatus36Choice)
	return s.Pending
}

func (s *SettlementStatus16Choice) AddFailing() *FailingStatus9Choice {
	s.Failing = new(FailingStatus9Choice)
	return s.Failing
}

func (s *SettlementStatus16Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	s.Proprietary = new(ProprietaryStatusAndReason6)
	return s.Proprietary
}
