package model

// Account to or from which a securities entry is made.
type SubAccountIdentification45 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount34 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation32 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification45) AddAccountOwner() *PartyIdentification109 {
	s.AccountOwner = new(PartyIdentification109)
	return s.AccountOwner
}

func (s *SubAccountIdentification45) AddSafekeepingAccount() *SecuritiesAccount34 {
	s.SafekeepingAccount = new(SecuritiesAccount34)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification45) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification45) AddBalanceForSubAccount() *AggregateBalanceInformation32 {
	newValue := new(AggregateBalanceInformation32)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}
