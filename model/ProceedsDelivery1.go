package model

// Provides information about the account.
type ProceedsDelivery1 struct {

	// identification of the securities account to which the securities have to be delivered.
	SecuritiesAccountIdentification *Max35Text `xml:"SctiesAcctId"`

	// Identification of the cash account to which the cash has to be delivered.
	CashAccountIdentification *CashAccountIdentification1Choice `xml:"CshAcctId"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the institution servicing the account.
	AccountServicerIdentification *PartyIdentification2Choice `xml:"AcctSvcrId,omitempty"`
}

func (p *ProceedsDelivery1) SetSecuritiesAccountIdentification(value string) {
	p.SecuritiesAccountIdentification = (*Max35Text)(&value)
}

func (p *ProceedsDelivery1) AddCashAccountIdentification() *CashAccountIdentification1Choice {
	p.CashAccountIdentification = new(CashAccountIdentification1Choice)
	return p.CashAccountIdentification
}

func (p *ProceedsDelivery1) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	p.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return p.AccountOwnerIdentification
}

func (p *ProceedsDelivery1) AddAccountServicerIdentification() *PartyIdentification2Choice {
	p.AccountServicerIdentification = new(PartyIdentification2Choice)
	return p.AccountServicerIdentification
}
