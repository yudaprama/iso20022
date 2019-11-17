package model

// Identification of an entity involved in an activity.
type PartyIdentification2 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification12Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentification2) AddIdentification() *PartyIdentification12Choice {
	p.Identification = new(PartyIdentification12Choice)
	return p.Identification
}

func (p *PartyIdentification2) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentification2) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification2) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification2) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}
