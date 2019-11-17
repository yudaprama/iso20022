package model

// Account to or from which a securities entry is made.
type SubAccountIdentification11 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation9 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification11) AddAccountOwner() *PartyIdentification13Choice {
	s.AccountOwner = new(PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification11) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification11) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification11) AddBalanceForSubAccount() *AggregateBalanceInformation9 {
	newValue := new(AggregateBalanceInformation9)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}
