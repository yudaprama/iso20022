package model

// Choice between a unique account identification and a set of account selection criteria.
type InvestmentAccountSelection2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Various investment account information used to select a specific account.
	OtherAccountSelectionData *InvestmentAccount29 `xml:"OthrAcctSelctnData"`
}

func (i *InvestmentAccountSelection2) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccountSelection2) AddOtherAccountSelectionData() *InvestmentAccount29 {
	i.OtherAccountSelectionData = new(InvestmentAccount29)
	return i.OtherAccountSelectionData
}
