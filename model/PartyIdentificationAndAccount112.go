package model

// Party and account details.
type PartyIdentificationAndAccount112 struct {

	// Identification of the party.
	Identification *PartyIdentification94Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

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

func (p *PartyIdentificationAndAccount112) AddIdentification() *PartyIdentification94Choice {
	p.Identification = new(PartyIdentification94Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount112) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount112) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount112) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount112) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount112) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount112) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount112) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}
