package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount60 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Account type.
	Type *InvestmentAccountType1Choice `xml:"Tp,omitempty"`
}

func (i *InvestmentAccount60) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount60) AddType() *InvestmentAccountType1Choice {
	i.Type = new(InvestmentAccountType1Choice)
	return i.Type
}
