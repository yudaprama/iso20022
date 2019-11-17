package model

// Identification of the party.
type PartyIdentification100 struct {

	// Unique identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification100) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentification100) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
