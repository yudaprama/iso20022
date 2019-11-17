package model

// Entity involved in an activity.
type PartyIdentificationAndAccount6 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	PartyIdentification *PartyIdentification25 `xml:"PtyId"`

	// Unambiguous identification of an account held by Financing Requestor to First Agent. This account is requested to be used for crediting the amount financed, as a result of the financing process.
	CreditAccount *CashAccount7 `xml:"CdtAcct,omitempty"`

	// Unambiguous identification of an internal bank account used by First Agent to manage the line of credit granted to Financing Requestor. This account is requested to be used for managing the financing process.
	FinancingAccount *CashAccount7 `xml:"FincgAcct,omitempty"`
}

func (p *PartyIdentificationAndAccount6) AddPartyIdentification() *PartyIdentification25 {
	p.PartyIdentification = new(PartyIdentification25)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount6) AddCreditAccount() *CashAccount7 {
	p.CreditAccount = new(CashAccount7)
	return p.CreditAccount
}

func (p *PartyIdentificationAndAccount6) AddFinancingAccount() *CashAccount7 {
	p.FinancingAccount = new(CashAccount7)
	return p.FinancingAccount
}
