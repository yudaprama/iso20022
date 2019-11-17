package model

// Party and account details.
type PartyIdentificationAndAccount20 struct {

	// Identification of the party.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount20) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount20) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount20) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount20) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount20) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount20) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount20) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}
