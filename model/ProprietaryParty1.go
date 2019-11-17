package model

// Set of elements to identify a proprietary party.
type ProprietaryParty1 struct {

	// Identifies the type of proprietary party reported.
	Type *Max35Text `xml:"Tp"`

	// Proprietary party.
	Party *PartyIdentification8 `xml:"Pty"`
}

func (p *ProprietaryParty1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryParty1) AddParty() *PartyIdentification8 {
	p.Party = new(PartyIdentification8)
	return p.Party
}
