package model

// Sub-account reporting.
type SubAccountIdentification34 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails17 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification34) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification34) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification34) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification34) AddFinancialInstrumentDetails() *FinancialInstrumentDetails17 {
	newValue := new(FinancialInstrumentDetails17)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}
