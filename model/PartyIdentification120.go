package model

// Provides information about identification of the party .
type PartyIdentification120 struct {

	// Identification of a party.
	Identification *PartyIdentification58Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification120) AddIdentification() *PartyIdentification58Choice {
	p.Identification = new(PartyIdentification58Choice)
	return p.Identification
}

func (p *PartyIdentification120) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentification120) AddAlternateIdentification() *AlternatePartyIdentification9 {
	newValue := new(AlternatePartyIdentification9)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}
