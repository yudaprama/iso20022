package model

// Party type and party identification information.
type PartyAndType1 struct {

	// Type of additional party.
	Type *PartyType1Choice `xml:"Tp"`

	// Details related to the additional party.
	Party *PartyIdentification43 `xml:"Pty,omitempty"`
}

func (p *PartyAndType1) AddType() *PartyType1Choice {
	p.Type = new(PartyType1Choice)
	return p.Type
}

func (p *PartyAndType1) AddParty() *PartyIdentification43 {
	p.Party = new(PartyIdentification43)
	return p.Party
}
