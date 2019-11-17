package model

// Choice of criteria for the identification of an account.
type AccountSelection1Choice struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Various investment account information used to select a specific account.
	OtherAccountSelectionData *InvestmentAccount52 `xml:"OthrAcctSelctnData"`
}

func (a *AccountSelection1Choice) SetAccountIdentification(value string) {
	a.AccountIdentification = (*Max35Text)(&value)
}

func (a *AccountSelection1Choice) AddOtherAccountSelectionData() *InvestmentAccount52 {
	a.OtherAccountSelectionData = new(InvestmentAccount52)
	return a.OtherAccountSelectionData
}
