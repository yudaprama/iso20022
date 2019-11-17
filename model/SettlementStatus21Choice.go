package model

// Choice of format for the settlement status.
type SettlementStatus21Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus47Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus12Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (s *SettlementStatus21Choice) AddPending() *PendingStatus47Choice {
	s.Pending = new(PendingStatus47Choice)
	return s.Pending
}

func (s *SettlementStatus21Choice) AddFailing() *FailingStatus12Choice {
	s.Failing = new(FailingStatus12Choice)
	return s.Failing
}

func (s *SettlementStatus21Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	s.Proprietary = new(ProprietaryStatusAndReason7)
	return s.Proprietary
}
