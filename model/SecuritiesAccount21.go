package model

// Unambiguous identification for the account between the account owner and the account servicer.
type SecuritiesAccount21 struct {

	// Account identification.
	Account *AccountIdentification5 `xml:"Acct"`

	// Sub-account identification.
	SubAccount *AccountIdentification5 `xml:"SubAcct,omitempty"`

	// Base currency for the account.
	BaseCurrency *ActiveOrHistoricCurrencyCode `xml:"BaseCcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReportingCurrency *ActiveOrHistoricCurrencyCode `xml:"RptgCcy,omitempty"`

	// Foreign exchange rate applied between the reporting and base currencies. It is assumed the valuation date is the same as the report date.
	ForeignExchangeRate *BaseOneRate `xml:"FXRate,omitempty"`
}

func (s *SecuritiesAccount21) AddAccount() *AccountIdentification5 {
	s.Account = new(AccountIdentification5)
	return s.Account
}

func (s *SecuritiesAccount21) AddSubAccount() *AccountIdentification5 {
	s.SubAccount = new(AccountIdentification5)
	return s.SubAccount
}

func (s *SecuritiesAccount21) SetBaseCurrency(value string) {
	s.BaseCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SecuritiesAccount21) SetReportingCurrency(value string) {
	s.ReportingCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SecuritiesAccount21) SetForeignExchangeRate(value string) {
	s.ForeignExchangeRate = (*BaseOneRate)(&value)
}
