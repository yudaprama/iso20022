package model

// Party and account details.
type PartyIdentificationAndAccount135 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id,omitempty"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Nationality of the investor or country of incorporation (for a company).
	Nationality *CountryCode `xml:"Ntlty,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount135) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount135) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount135) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount135) SetNationality(value string) {
	p.Nationality = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount135) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (p *PartyIdentificationAndAccount135) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount135) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}
