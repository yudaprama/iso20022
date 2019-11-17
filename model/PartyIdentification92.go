package model

// Provides information about identification of the party .
type PartyIdentification92 struct {

	// Identification of a party.
	Identification *PartyIdentification44Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification92) AddIdentification() *PartyIdentification44Choice {
	p.Identification = new(PartyIdentification44Choice)
	return p.Identification
}

func (p *PartyIdentification92) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification92) AddAlternateIdentification() *AlternatePartyIdentification7 {
	newValue := new(AlternatePartyIdentification7)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}
