package model

// Identification of an entity involved in an activity.
type PartyIdentification55 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification68Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentification55) AddIdentification() *PartyIdentification68Choice {
	p.Identification = new(PartyIdentification68Choice)
	return p.Identification
}

func (p *PartyIdentification55) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentification55) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}
