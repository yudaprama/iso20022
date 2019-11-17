package model

// Party and account details.
type PartyIdentificationAndAccount127 struct {

	// Identification of the party.
	Identification *PartyIdentification101Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification8 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount127) AddIdentification() *PartyIdentification101Choice {
	p.Identification = new(PartyIdentification101Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount127) AddAlternateIdentification() *AlternatePartyIdentification8 {
	p.AlternateIdentification = new(AlternatePartyIdentification8)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount127) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount127) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount127) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}
