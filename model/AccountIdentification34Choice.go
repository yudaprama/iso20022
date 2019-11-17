package model

// Choice between all accounts (GENR - General in ISO 15022) or one or more selected accounts.
type AccountIdentification34Choice struct {

	// All safekeeping accounts that own underlying financial instrument.
	ForAllAccounts *AccountIdentification10 `xml:"ForAllAccts"`

	// Selected safekeeping accounts list to which the corporate action event applies.
	AccountsList []*AccountIdentification34 `xml:"AcctsList"`
}

func (a *AccountIdentification34Choice) AddForAllAccounts() *AccountIdentification10 {
	a.ForAllAccounts = new(AccountIdentification10)
	return a.ForAllAccounts
}

func (a *AccountIdentification34Choice) AddAccountsList() *AccountIdentification34 {
	newValue := new(AccountIdentification34)
	a.AccountsList = append(a.AccountsList, newValue)
	return newValue
}
