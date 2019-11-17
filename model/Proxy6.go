package model

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy6 struct {

	// Specifies the type of proxy.
	ProxyType *ProxyType2Code `xml:"PrxyTp"`

	// Person, other than the chairman of the meeting, assigned by the security holder as the proxy.
	PersonDetails *IndividualPerson26 `xml:"PrsnDtls,omitempty"`

	// Indicates the vote instruction for the resolutions that are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote3Choice `xml:"VoteInstrForAgndRsltn,omitempty"`
}

func (p *Proxy6) SetProxyType(value string) {
	p.ProxyType = (*ProxyType2Code)(&value)
}

func (p *Proxy6) AddPersonDetails() *IndividualPerson26 {
	p.PersonDetails = new(IndividualPerson26)
	return p.PersonDetails
}

func (p *Proxy6) AddVoteInstructionForAgendaResolution() *Vote3Choice {
	p.VoteInstructionForAgendaResolution = new(Vote3Choice)
	return p.VoteInstructionForAgendaResolution
}
