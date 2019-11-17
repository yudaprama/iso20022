package model

// Provides information about identification of the party .
type PartyIdentification102 struct {

	// Identification of a party.
	Identification *PartyIdentification111Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification102) AddIdentification() *PartyIdentification111Choice {
	p.Identification = new(PartyIdentification111Choice)
	return p.Identification
}

func (p *PartyIdentification102) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentification102) AddAlternateIdentification() *AlternatePartyIdentification9 {
	newValue := new(AlternatePartyIdentification9)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}
