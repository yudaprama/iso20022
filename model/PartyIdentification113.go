package model

// Identification of a party.
type PartyIdentification113 struct {

	// Unique identification of the party.
	Party *PartyIdentification90Choice `xml:"Pty"`

	// Legal entity identification as an alternate identification for the party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification113) AddParty() *PartyIdentification90Choice {
	p.Party = new(PartyIdentification90Choice)
	return p.Party
}

func (p *PartyIdentification113) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
