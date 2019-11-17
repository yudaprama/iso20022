package model

// Identifies the securities for which the meeting is organised.
type SecurityPosition5 struct {

	// Security held in an account on which the balance is calculated.
	Identification *SecurityIdentification3 `xml:"Id"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition2 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition5) AddIdentification() *SecurityIdentification3 {
	s.Identification = new(SecurityIdentification3)
	return s.Identification
}

func (s *SecurityPosition5) AddPosition() *EligiblePosition2 {
	newValue := new(EligiblePosition2)
	s.Position = append(s.Position, newValue)
	return newValue
}
