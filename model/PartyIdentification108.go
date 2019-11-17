package model

// Identification of an entity involved in an activity.
type PartyIdentification108 struct {

	// Identification of the party.
	Identification *PartyIdentification58Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification108) AddIdentification() *PartyIdentification58Choice {
	p.Identification = new(PartyIdentification58Choice)
	return p.Identification
}

func (p *PartyIdentification108) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentification108) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}
