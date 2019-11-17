package model

// Identification of the party.
type PartyIdentification119 struct {

	// Unique identification of the party.
	Identification *PartyIdentification103Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification119) AddIdentification() *PartyIdentification103Choice {
	p.Identification = new(PartyIdentification103Choice)
	return p.Identification
}

func (p *PartyIdentification119) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
