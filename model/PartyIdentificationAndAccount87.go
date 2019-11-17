package model

// Entity involved in an activity.
type PartyIdentificationAndAccount87 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

	// Entity involved in an activity.
	AlternateIdentification *AlternatePartyIdentification6 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount87) AddIdentification() *PartyIdentification70Choice {
	p.Identification = new(PartyIdentification70Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount87) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount87) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

func (p *PartyIdentificationAndAccount87) AddAlternateIdentification() *AlternatePartyIdentification6 {
	p.AlternateIdentification = new(AlternatePartyIdentification6)
	return p.AlternateIdentification
}
