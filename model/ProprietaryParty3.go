package model

// Identifies a proprietary party.
type ProprietaryParty3 struct {

	// Specifies the type of proprietary party.
	Type *Max35Text `xml:"Tp"`

	// Proprietary party.
	Party *PartyIdentification43 `xml:"Pty"`
}

func (p *ProprietaryParty3) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryParty3) AddParty() *PartyIdentification43 {
	p.Party = new(PartyIdentification43)
	return p.Party
}
