package model

// Party and account details.
type PartyIdentificationAndAccount78 struct {

	// Identification of the party.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Coded list to specify the side of the order.
	Side *ClearingSide1Code `xml:"Sd,omitempty"`

	// Identifies the clearing member account at the CCP through which the trade must be cleared (sometimes called position account).
	ClearingAccount *SecuritiesAccount20 `xml:"ClrAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount78) AddIdentification() *PartyIdentification32Choice {
	p.Identification = new(PartyIdentification32Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount78) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount78) SetSide(value string) {
	p.Side = (*ClearingSide1Code)(&value)
}

func (p *PartyIdentificationAndAccount78) AddClearingAccount() *SecuritiesAccount20 {
	p.ClearingAccount = new(SecuritiesAccount20)
	return p.ClearingAccount
}

func (p *PartyIdentificationAndAccount78) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount78) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}
