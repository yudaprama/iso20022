package model

// Choice of format for the settlement status.
type SettlementStatus20Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus45Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus11Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (s *SettlementStatus20Choice) AddPending() *PendingStatus45Choice {
	s.Pending = new(PendingStatus45Choice)
	return s.Pending
}

func (s *SettlementStatus20Choice) AddFailing() *FailingStatus11Choice {
	s.Failing = new(FailingStatus11Choice)
	return s.Failing
}

func (s *SettlementStatus20Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	s.Proprietary = new(ProprietaryStatusAndReason7)
	return s.Proprietary
}
