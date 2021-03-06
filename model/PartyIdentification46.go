package model

// Identification of an entity involved in an activity.
type PartyIdentification46 struct {

	// Identification of the party.
	Identification *PartyIdentification44Choice `xml:"Id"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification46) AddIdentification() *PartyIdentification44Choice {
	p.Identification = new(PartyIdentification44Choice)
	return p.Identification
}

func (p *PartyIdentification46) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}
