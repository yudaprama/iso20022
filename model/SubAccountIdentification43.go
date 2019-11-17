package model

// Account to or from which a securities entry is made.
type SubAccountIdentification43 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount25 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation31 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification43) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SubAccountIdentification43) AddSafekeepingAccount() *SecuritiesAccount25 {
	s.SafekeepingAccount = new(SecuritiesAccount25)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification43) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification43) AddBalanceForSubAccount() *AggregateBalanceInformation31 {
	newValue := new(AggregateBalanceInformation31)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}
