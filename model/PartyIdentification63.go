package model

// Identification of an entity involved in an activity.
type PartyIdentification63 struct {

	// Identification of the party.
	PartyIdentification *PartyIdentification75Choice `xml:"PtyId"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification63) AddPartyIdentification() *PartyIdentification75Choice {
	p.PartyIdentification = new(PartyIdentification75Choice)
	return p.PartyIdentification
}

func (p *PartyIdentification63) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}
