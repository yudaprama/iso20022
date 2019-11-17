package model

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount129 struct {

	// Identification of a party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount129) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount129) AddCashAccount() *CashAccountIdentification6Choice {
	p.CashAccount = new(CashAccountIdentification6Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount129) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount129) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}
