package model

// Identification of the party.
type PartyIdentification111 struct {

	// Unique identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification111) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentification111) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
