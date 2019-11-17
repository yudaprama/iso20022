package model

// Entity involved in an activity.
type PartyIdentificationAndAccount79 struct {

	// Information related to an identification, eg, party identification or account identification.
	Identification *PartyIdentification32Choice `xml:"Id,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification2Choice `xml:"CshAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount79) AddIdentification() *PartyIdentification32Choice {
	p.Identification = new(PartyIdentification32Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount79) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount79) AddCashAccount() *CashAccountIdentification2Choice {
	p.CashAccount = new(CashAccountIdentification2Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount79) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount79) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount79) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

func (p *PartyIdentificationAndAccount79) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}
