package model

// Specifies a choice between one party or a group of parties.
type PartyOrGroup1Choice struct {

	// Specifies the identification of a group of parties.
	GroupIdentification *Max4AlphaNumericText `xml:"GrpId"`

	// Specifies a party.
	Party *PartyAndCertificate2 `xml:"Pty"`
}

func (p *PartyOrGroup1Choice) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max4AlphaNumericText)(&value)
}

func (p *PartyOrGroup1Choice) AddParty() *PartyAndCertificate2 {
	p.Party = new(PartyAndCertificate2)
	return p.Party
}
