package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount15 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Account type.
	Type *CashAccountType1 `xml:"Tp,omitempty"`
}

func (i *InvestmentAccount15) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount15) AddType() *CashAccountType1 {
	i.Type = new(CashAccountType1)
	return i.Type
}
