package model

// Provides all sub-account details.
type SubAccountIdentification44 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount25 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails20 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification44) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SubAccountIdentification44) AddSafekeepingAccount() *SecuritiesAccount25 {
	s.SafekeepingAccount = new(SecuritiesAccount25)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification44) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification44) AddFinancialInstrumentDetails() *FinancialInstrumentDetails20 {
	newValue := new(FinancialInstrumentDetails20)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}
