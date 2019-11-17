package model

// Choice of criteria for the identification of an account.
type AccountSelection2Choice struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Various investment account information used to select a specific account.
	OtherAccountSelectionData *InvestmentAccount64 `xml:"OthrAcctSelctnData"`
}

func (a *AccountSelection2Choice) SetAccountIdentification(value string) {
	a.AccountIdentification = (*Max35Text)(&value)
}

func (a *AccountSelection2Choice) AddOtherAccountSelectionData() *InvestmentAccount64 {
	a.OtherAccountSelectionData = new(InvestmentAccount64)
	return a.OtherAccountSelectionData
}
