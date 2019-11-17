package model

// Identifies the securities for which the meeting is organised.
type SecurityPosition7 struct {

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification11 `xml:"Id"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition4 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition7) AddIdentification() *SecurityIdentification11 {
	s.Identification = new(SecurityIdentification11)
	return s.Identification
}

func (s *SecurityPosition7) AddPosition() *EligiblePosition4 {
	newValue := new(EligiblePosition4)
	s.Position = append(s.Position, newValue)
	return newValue
}
