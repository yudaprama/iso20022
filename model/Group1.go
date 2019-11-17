package model

// Group of parties with their related security certificate.
type Group1 struct {

	// Specifies the identification of the group.
	GroupIdentification *Max4AlphaNumericText `xml:"GrpId"`

	// Specifies a party and related certificate.
	Party []*PartyAndCertificate2 `xml:"Pty"`
}

func (g *Group1) SetGroupIdentification(value string) {
	g.GroupIdentification = (*Max4AlphaNumericText)(&value)
}

func (g *Group1) AddParty() *PartyAndCertificate2 {
	newValue := new(PartyAndCertificate2)
	g.Party = append(g.Party, newValue)
	return newValue
}
