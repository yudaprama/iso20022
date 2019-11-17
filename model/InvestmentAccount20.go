package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount20 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Account type.
	Type *FundCashAccount2Code `xml:"Tp,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`
}

func (i *InvestmentAccount20) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount20) SetType(value string) {
	i.Type = (*FundCashAccount2Code)(&value)
}

func (i *InvestmentAccount20) SetExtendedType(value string) {
	i.ExtendedType = (*Extended350Code)(&value)
}
