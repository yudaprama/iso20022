package model

// Sub-account reporting.
type SubAccountIdentification21 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails9 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification21) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification21) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification21) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification21) AddFinancialInstrumentDetails() *FinancialInstrumentDetails9 {
	newValue := new(FinancialInstrumentDetails9)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}
