package model

// Request to execute specific instructions, such as participation registration, securities registration or blocking of securities.
type SpecificInstructionRequest1 struct {

	// Request to register for participation to the meeting.
	ParticipationRegistration *YesNoIndicator `xml:"PrtcptnRegn,omitempty"`

	// Request to block the securities
	BlockingSecurities *YesNoIndicator `xml:"BlckgScties,omitempty"`

	// Request to register the securities for the meeting.
	SecuritiesRegistration *YesNoIndicator `xml:"SctiesRegn,omitempty"`
}

func (s *SpecificInstructionRequest1) SetParticipationRegistration(value string) {
	s.ParticipationRegistration = (*YesNoIndicator)(&value)
}

func (s *SpecificInstructionRequest1) SetBlockingSecurities(value string) {
	s.BlockingSecurities = (*YesNoIndicator)(&value)
}

func (s *SpecificInstructionRequest1) SetSecuritiesRegistration(value string) {
	s.SecuritiesRegistration = (*YesNoIndicator)(&value)
}
