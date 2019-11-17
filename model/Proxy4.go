package model

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy4 struct {

	// Specifies the type of proxy.
	ProxyType *ProxyType2Code `xml:"PrxyTp"`

	// Person, other than the Chairman of the meeting, assigned by the security holder as proxy.
	PersonDetails *IndividualPerson17 `xml:"PrsnDtls,omitempty"`

	// Indicates the vote instruction for the resolutions which are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote2Choice `xml:"VoteInstrForAgndRsltn,omitempty"`
}

func (p *Proxy4) SetProxyType(value string) {
	p.ProxyType = (*ProxyType2Code)(&value)
}

func (p *Proxy4) AddPersonDetails() *IndividualPerson17 {
	p.PersonDetails = new(IndividualPerson17)
	return p.PersonDetails
}

func (p *Proxy4) AddVoteInstructionForAgendaResolution() *Vote2Choice {
	p.VoteInstructionForAgendaResolution = new(Vote2Choice)
	return p.VoteInstructionForAgendaResolution
}
