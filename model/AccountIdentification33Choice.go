package model

// Choice between all accounts (GENR - General in ISO 15022) or one or more selected accounts and balance information.
type AccountIdentification33Choice struct {

	// All safekeeping accounts that own underlying financial instrument.
	ForAllAccounts *AccountIdentification10 `xml:"ForAllAccts"`

	// Selected safekeeping accounts list (and optionally balance information) to which the corporate action event applies.
	AccountsListAndBalanceDetails []*AccountIdentification32 `xml:"AcctsListAndBalDtls"`
}

func (a *AccountIdentification33Choice) AddForAllAccounts() *AccountIdentification10 {
	a.ForAllAccounts = new(AccountIdentification10)
	return a.ForAllAccounts
}

func (a *AccountIdentification33Choice) AddAccountsListAndBalanceDetails() *AccountIdentification32 {
	newValue := new(AccountIdentification32)
	a.AccountsListAndBalanceDetails = append(a.AccountsListAndBalanceDetails, newValue)
	return newValue
}
