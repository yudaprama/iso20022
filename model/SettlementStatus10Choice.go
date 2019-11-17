package model

// Choice of format for the settlement status.
type SettlementStatus10Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus15Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus3Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (s *SettlementStatus10Choice) AddPending() *PendingStatus15Choice {
	s.Pending = new(PendingStatus15Choice)
	return s.Pending
}

func (s *SettlementStatus10Choice) AddFailing() *FailingStatus3Choice {
	s.Failing = new(FailingStatus3Choice)
	return s.Failing
}

func (s *SettlementStatus10Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	s.Proprietary = new(ProprietaryStatusAndReason1)
	return s.Proprietary
}
