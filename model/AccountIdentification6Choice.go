package model

// Choice between all accounts (GENR - General in ISO 15022) or one or more selected accounts.
type AccountIdentification6Choice struct {

	// All safekeeping accounts that own underlying financial instrument.
	ForAllAccounts *AccountIdentification10 `xml:"ForAllAccts"`

	// Selected safekeeping accounts list to which the corporate action event applies.
	AccountsList []*AccountIdentification9 `xml:"AcctsList"`
}

func (a *AccountIdentification6Choice) AddForAllAccounts() *AccountIdentification10 {
	a.ForAllAccounts = new(AccountIdentification10)
	return a.ForAllAccounts
}

func (a *AccountIdentification6Choice) AddAccountsList() *AccountIdentification9 {
	newValue := new(AccountIdentification9)
	a.AccountsList = append(a.AccountsList, newValue)
	return newValue
}
