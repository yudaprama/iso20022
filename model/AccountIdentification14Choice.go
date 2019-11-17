package model

// Choice between all accounts (GENR - General in ISO 15022) or one or more selected accounts and balance information.
type AccountIdentification14Choice struct {

	// All safekeeping accounts that own underlying financial instrument.
	ForAllAccounts *AccountIdentification10 `xml:"ForAllAccts"`

	// Selected safekeeping accounts list (and optionally balance information) to which the corporate action event applies.
	AccountsListAndBalanceDetails []*AccountIdentification18 `xml:"AcctsListAndBalDtls"`
}

func (a *AccountIdentification14Choice) AddForAllAccounts() *AccountIdentification10 {
	a.ForAllAccounts = new(AccountIdentification10)
	return a.ForAllAccounts
}

func (a *AccountIdentification14Choice) AddAccountsListAndBalanceDetails() *AccountIdentification18 {
	newValue := new(AccountIdentification18)
	a.AccountsListAndBalanceDetails = append(a.AccountsListAndBalanceDetails, newValue)
	return newValue
}
