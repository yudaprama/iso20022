package model

// Choice of format for the settlement status.
type SettlementStatus9Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus12Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus4Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (s *SettlementStatus9Choice) AddPending() *PendingStatus12Choice {
	s.Pending = new(PendingStatus12Choice)
	return s.Pending
}

func (s *SettlementStatus9Choice) AddFailing() *FailingStatus4Choice {
	s.Failing = new(FailingStatus4Choice)
	return s.Failing
}

func (s *SettlementStatus9Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	s.Proprietary = new(ProprietaryStatusAndReason1)
	return s.Proprietary
}
