package model

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount54 struct {

	// Identification of a party.
	Identification *PartyIdentification42Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount54) AddIdentification() *PartyIdentification42Choice {
	p.Identification = new(PartyIdentification42Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount54) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount54) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount54) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount54) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}
