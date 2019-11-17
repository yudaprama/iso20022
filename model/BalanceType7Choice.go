package model

// Choice of format for the balance type.
type BalanceType7Choice struct {

	// Balance type expressed in coded form.
	Code *FinancialAssetBalanceType1Code `xml:"Cd"`

	// Balance type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`

	// Account identification.
	Account *AccountIdentification5 `xml:"Acct"`
}

func (b *BalanceType7Choice) SetCode(value string) {
	b.Code = (*FinancialAssetBalanceType1Code)(&value)
}

func (b *BalanceType7Choice) AddProprietary() *GenericIdentification30 {
	b.Proprietary = new(GenericIdentification30)
	return b.Proprietary
}

func (b *BalanceType7Choice) AddAccount() *AccountIdentification5 {
	b.Account = new(AccountIdentification5)
	return b.Account
}
