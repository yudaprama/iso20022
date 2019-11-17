package model

// Provides all sub-account details.
type SubAccountIdentification50 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount34 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails27 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification50) AddAccountOwner() *PartyIdentification119 {
	s.AccountOwner = new(PartyIdentification119)
	return s.AccountOwner
}

func (s *SubAccountIdentification50) AddSafekeepingAccount() *SecuritiesAccount34 {
	s.SafekeepingAccount = new(SecuritiesAccount34)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification50) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification50) AddFinancialInstrumentDetails() *FinancialInstrumentDetails27 {
	newValue := new(FinancialInstrumentDetails27)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}
