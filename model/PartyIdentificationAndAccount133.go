package model

// Party and account details.
type PartyIdentificationAndAccount133 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification6Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification6Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification6Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation4 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount133) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount133) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount133) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount133) AddCashAccount() *CashAccountIdentification6Choice {
	p.CashAccount = new(CashAccountIdentification6Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount133) AddChargesAccount() *CashAccountIdentification6Choice {
	p.ChargesAccount = new(CashAccountIdentification6Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount133) AddCommissionAccount() *CashAccountIdentification6Choice {
	p.CommissionAccount = new(CashAccountIdentification6Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount133) AddTaxAccount() *CashAccountIdentification6Choice {
	p.TaxAccount = new(CashAccountIdentification6Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount133) AddAdditionalInformation() *PartyTextInformation4 {
	p.AdditionalInformation = new(PartyTextInformation4)
	return p.AdditionalInformation
}
