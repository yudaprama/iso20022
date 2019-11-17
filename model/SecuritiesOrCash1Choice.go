package model

// Choice of securities or cash parties.
type SecuritiesOrCash1Choice struct {

	// Securities settlement chain parties, accounts and other details.
	SecuritiesDetails *SettlementParties35 `xml:"SctiesDtls"`

	// Cash settlement chain parties and accounts.
	CashPartiesDetails *CashParties24 `xml:"CshPtiesDtls"`
}

func (s *SecuritiesOrCash1Choice) AddSecuritiesDetails() *SettlementParties35 {
	s.SecuritiesDetails = new(SettlementParties35)
	return s.SecuritiesDetails
}

func (s *SecuritiesOrCash1Choice) AddCashPartiesDetails() *CashParties24 {
	s.CashPartiesDetails = new(CashParties24)
	return s.CashPartiesDetails
}
