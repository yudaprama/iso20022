package model

// Identifies the securities for which the meeting is organised.
type SecurityPosition6 struct {

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification11 `xml:"Id"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition3 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition6) AddIdentification() *SecurityIdentification11 {
	s.Identification = new(SecurityIdentification11)
	return s.Identification
}

func (s *SecurityPosition6) AddPosition() *EligiblePosition3 {
	newValue := new(EligiblePosition3)
	s.Position = append(s.Position, newValue)
	return newValue
}
