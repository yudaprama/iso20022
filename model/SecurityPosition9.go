package model

// Identifies the securities for which the meeting is organised.
type SecurityPosition9 struct {

	// Identification of the security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition6 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition9) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecurityPosition9) AddPosition() *EligiblePosition6 {
	newValue := new(EligiblePosition6)
	s.Position = append(s.Position, newValue)
	return newValue
}
