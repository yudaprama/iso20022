package model

// Party and account details.
type PartyIdentificationAndAccount137 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount137) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount137) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount137) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount137) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount137) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}
