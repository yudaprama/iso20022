package model

// Identification of the party.
type PartyIdentification109 struct {

	// Unique identification of the party.
	Identification *PartyIdentification114Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification109) AddIdentification() *PartyIdentification114Choice {
	p.Identification = new(PartyIdentification114Choice)
	return p.Identification
}

func (p *PartyIdentification109) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
