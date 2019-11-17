package model

// Identification of the party.
type PartyIdentification98 struct {

	// Unique identification of the party.
	Identification *PartyIdentification92Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification98) AddIdentification() *PartyIdentification92Choice {
	p.Identification = new(PartyIdentification92Choice)
	return p.Identification
}

func (p *PartyIdentification98) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
