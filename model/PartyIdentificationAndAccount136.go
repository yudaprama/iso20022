package model

// Party and account details.
type PartyIdentificationAndAccount136 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount136) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount136) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount136) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount136) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (p *PartyIdentificationAndAccount136) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount136) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}
