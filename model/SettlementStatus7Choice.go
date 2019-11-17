package model

// Choice of format for the settlement status.
type SettlementStatus7Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus9Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus3Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (s *SettlementStatus7Choice) AddPending() *PendingStatus9Choice {
	s.Pending = new(PendingStatus9Choice)
	return s.Pending
}

func (s *SettlementStatus7Choice) AddFailing() *FailingStatus3Choice {
	s.Failing = new(FailingStatus3Choice)
	return s.Failing
}

func (s *SettlementStatus7Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	s.Proprietary = new(ProprietaryStatusAndReason1)
	return s.Proprietary
}
