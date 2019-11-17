package model

// Set of elements used to identify a proprietary party.
type ProprietaryParty2 struct {

	// Specifies the type of proprietary party.
	Type *Max35Text `xml:"Tp"`

	// Proprietary party.
	Party *PartyIdentification32 `xml:"Pty"`
}

func (p *ProprietaryParty2) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryParty2) AddParty() *PartyIdentification32 {
	p.Party = new(PartyIdentification32)
	return p.Party
}
