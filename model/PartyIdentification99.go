package model

// Identification of the party.
type PartyIdentification99 struct {

	// Unique identification of the party.
	Identification *PartyIdentification93Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification99) AddIdentification() *PartyIdentification93Choice {
	p.Identification = new(PartyIdentification93Choice)
	return p.Identification
}

func (p *PartyIdentification99) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
