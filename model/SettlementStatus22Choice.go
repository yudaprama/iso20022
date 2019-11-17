package model

// Choice of format for the settlement status.
type SettlementStatus22Choice struct {

	// Instruction is pending. Settlement at the instructed settlement date is still possible.
	Pending *PendingStatus50Choice `xml:"Pdg"`

	// Instruction is failing. Settlement at the instructed settlement date is no longer possible.
	Failing *FailingStatus12Choice `xml:"Flng"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (s *SettlementStatus22Choice) AddPending() *PendingStatus50Choice {
	s.Pending = new(PendingStatus50Choice)
	return s.Pending
}

func (s *SettlementStatus22Choice) AddFailing() *FailingStatus12Choice {
	s.Failing = new(FailingStatus12Choice)
	return s.Failing
}

func (s *SettlementStatus22Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	s.Proprietary = new(ProprietaryStatusAndReason7)
	return s.Proprietary
}
