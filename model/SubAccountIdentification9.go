package model

// Sub-account reporting.
type SubAccountIdentification9 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails2 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification9) AddAccountOwner() *PartyIdentification13Choice {
	s.AccountOwner = new(PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification9) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification9) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification9) AddFinancialInstrumentDetails() *FinancialInstrumentDetails2 {
	newValue := new(FinancialInstrumentDetails2)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}
