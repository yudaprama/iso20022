package model

// Provides all sub-account details.
type SubAccountIdentification47 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount34 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails23 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification47) AddAccountOwner() *PartyIdentification109 {
	s.AccountOwner = new(PartyIdentification109)
	return s.AccountOwner
}

func (s *SubAccountIdentification47) AddSafekeepingAccount() *SecuritiesAccount34 {
	s.SafekeepingAccount = new(SecuritiesAccount34)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification47) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification47) AddFinancialInstrumentDetails() *FinancialInstrumentDetails23 {
	newValue := new(FinancialInstrumentDetails23)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}
