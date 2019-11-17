package model

// Identifies the securities for which the meeting is organised.
type SecurityPosition8 struct {

	// Identification of the security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition5 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition8) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecurityPosition8) AddPosition() *EligiblePosition5 {
	newValue := new(EligiblePosition5)
	s.Position = append(s.Position, newValue)
	return newValue
}
