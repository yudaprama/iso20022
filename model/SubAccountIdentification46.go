package model

// Account to or from which a securities entry is made.
type SubAccountIdentification46 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount34 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation33 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification46) AddAccountOwner() *PartyIdentification109 {
	s.AccountOwner = new(PartyIdentification109)
	return s.AccountOwner
}

func (s *SubAccountIdentification46) AddSafekeepingAccount() *SecuritiesAccount34 {
	s.SafekeepingAccount = new(SecuritiesAccount34)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification46) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification46) AddBalanceForSubAccount() *AggregateBalanceInformation33 {
	newValue := new(AggregateBalanceInformation33)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}
