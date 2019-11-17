package model

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy2 struct {

	// Specifies the type of proxy.
	ProxyType *ProxyType2Code `xml:"PrxyTp"`

	// Person, other than the Chairman of the meeting, assigned by the security holder as proxy.
	PersonDetails *IndividualPerson13 `xml:"PrsnDtls,omitempty"`

	// Indicates the vote instruction for the resolutions which are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote1Choice `xml:"VoteInstrForAgndRsltn,omitempty"`
}

func (p *Proxy2) SetProxyType(value string) {
	p.ProxyType = (*ProxyType2Code)(&value)
}

func (p *Proxy2) AddPersonDetails() *IndividualPerson13 {
	p.PersonDetails = new(IndividualPerson13)
	return p.PersonDetails
}

func (p *Proxy2) AddVoteInstructionForAgendaResolution() *Vote1Choice {
	p.VoteInstructionForAgendaResolution = new(Vote1Choice)
	return p.VoteInstructionForAgendaResolution
}
