package model

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount121 struct {

	// Identification of a party.
	Identification *PartyIdentification94Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount121) AddIdentification() *PartyIdentification94Choice {
	p.Identification = new(PartyIdentification94Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount121) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount121) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount121) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}
