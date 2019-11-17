package model

// Account to or from which a securities entry is made.
type SubAccountIdentification10 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation8 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification10) AddAccountOwner() *PartyIdentification13Choice {
	s.AccountOwner = new(PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification10) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification10) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification10) AddBalanceForSubAccount() *AggregateBalanceInformation8 {
	newValue := new(AggregateBalanceInformation8)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}
