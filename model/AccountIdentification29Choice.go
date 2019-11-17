package model

// Choice between all accounts (GENR - General in ISO 15022) or one or more selected accounts.
type AccountIdentification29Choice struct {

	// All safekeeping accounts that own underlying financial instrument.
	ForAllAccounts *AccountIdentification10 `xml:"ForAllAccts"`

	// Selected safekeeping accounts list to which the corporate action event applies.
	AccountsList []*AccountIdentification31 `xml:"AcctsList"`
}

func (a *AccountIdentification29Choice) AddForAllAccounts() *AccountIdentification10 {
	a.ForAllAccounts = new(AccountIdentification10)
	return a.ForAllAccounts
}

func (a *AccountIdentification29Choice) AddAccountsList() *AccountIdentification31 {
	newValue := new(AccountIdentification31)
	a.AccountsList = append(a.AccountsList, newValue)
	return newValue
}
