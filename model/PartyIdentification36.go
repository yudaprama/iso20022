package model

// Identification of an entity involved in an activity.
type PartyIdentification36 struct {

	// Identification of the party.
	Identification *PartyIdentification12Choice `xml:"Id"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification36) AddIdentification() *PartyIdentification12Choice {
	p.Identification = new(PartyIdentification12Choice)
	return p.Identification
}

func (p *PartyIdentification36) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}
