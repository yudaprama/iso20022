package model

// Identification of the party.
type PartyIdentification110 struct {

	// Unique identification of the party.
	Identification *PartyIdentification115Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification110) AddIdentification() *PartyIdentification115Choice {
	p.Identification = new(PartyIdentification115Choice)
	return p.Identification
}

func (p *PartyIdentification110) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
