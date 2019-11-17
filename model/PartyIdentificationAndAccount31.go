package model

// Party and account details.
type PartyIdentificationAndAccount31 struct {

	// Identification of the party.
	Identification *PartyIdentification33Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification4 `xml:"AltrnId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

	// Identifies the clearing member account at the Central counterparty through which the trade must be cleared (sometimes called position account).
	ClearingAccount *SecuritiesAccount18 `xml:"ClrAcct,omitempty"`
}

func (p *PartyIdentificationAndAccount31) AddIdentification() *PartyIdentification33Choice {
	p.Identification = new(PartyIdentification33Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount31) AddAlternateIdentification() *AlternatePartyIdentification4 {
	p.AlternateIdentification = new(AlternatePartyIdentification4)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount31) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

func (p *PartyIdentificationAndAccount31) AddClearingAccount() *SecuritiesAccount18 {
	p.ClearingAccount = new(SecuritiesAccount18)
	return p.ClearingAccount
}
