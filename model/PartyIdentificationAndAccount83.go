package model

// Entity involved in an activity.
type PartyIdentificationAndAccount83 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification2Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

	// Entity involved in an activity.
	AlternateIdentification *AlternatePartyIdentification6 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount83) AddIdentification() *PartyIdentification70Choice {
	p.Identification = new(PartyIdentification70Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount83) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount83) AddCashAccount() *CashAccountIdentification2Choice {
	p.CashAccount = new(CashAccountIdentification2Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount83) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount83) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

func (p *PartyIdentificationAndAccount83) AddAlternateIdentification() *AlternatePartyIdentification6 {
	p.AlternateIdentification = new(AlternatePartyIdentification6)
	return p.AlternateIdentification
}
