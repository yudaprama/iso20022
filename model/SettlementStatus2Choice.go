package model

// Choice of format for the settlement status.
type SettlementStatus2Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus3Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus1Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (s *SettlementStatus2Choice) AddPending() *PendingStatus3Choice {
	s.Pending = new(PendingStatus3Choice)
	return s.Pending
}

func (s *SettlementStatus2Choice) AddFailing() *FailingStatus1Choice {
	s.Failing = new(FailingStatus1Choice)
	return s.Failing
}

func (s *SettlementStatus2Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	s.Proprietary = new(ProprietaryStatusAndReason1)
	return s.Proprietary
}
