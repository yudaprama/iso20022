package model

// Statement information of an account.
type ATMAccountStatement1 struct {

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Statement information.
	AccountStatement []*ATMAccountStatement2 `xml:"AcctStmt,omitempty"`
}

func (a *ATMAccountStatement1) AddAccountIdentifier() *AccountIdentification31Choice {
	a.AccountIdentifier = new(AccountIdentification31Choice)
	return a.AccountIdentifier
}

func (a *ATMAccountStatement1) SetAccountName(value string) {
	a.AccountName = (*Max70Text)(&value)
}

func (a *ATMAccountStatement1) AddAccountStatement() *ATMAccountStatement2 {
	newValue := new(ATMAccountStatement2)
	a.AccountStatement = append(a.AccountStatement, newValue)
	return newValue
}
