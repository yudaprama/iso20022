package model

// Identification of an entity involved in an activity.
type PartyIdentification75 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification44Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentification75) AddIdentification() *PartyIdentification44Choice {
	p.Identification = new(PartyIdentification44Choice)
	return p.Identification
}

func (p *PartyIdentification75) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentification75) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentification75) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification75) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification75) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}
