package model

// Specifies the cash-in and cash-out flows by a user defined parameter/s.
type BreakdownByUserDefinedParameter3 struct {

	// Party for which the cash flow is being reported.
	Party *InvestmentAccount42 `xml:"Pty,omitempty"`

	// Country for which the cash flow is being reported.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Currency for which the cash flow is being reported.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Parameter for which the cash flow is being reported.
	UserDefined *DataFormat2Choice `xml:"UsrDfnd,omitempty"`

	// Cash movement into the fund as a result of transactions in shares in an investment fund, for example, subscriptions or switch-ins.
	CashInForecast []*CashInForecast5 `xml:"CshInFcst,omitempty"`

	// Cash movement out of the fund as a result of transactions in shares in an investment fund, for example, redemptions or switch-outs.
	CashOutForecast []*CashOutForecast5 `xml:"CshOutFcst,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows specified for the user defined parameter.
	NetCashForecast []*NetCashForecast4 `xml:"NetCshFcst,omitempty"`
}

func (b *BreakdownByUserDefinedParameter3) AddParty() *InvestmentAccount42 {
	b.Party = new(InvestmentAccount42)
	return b.Party
}

func (b *BreakdownByUserDefinedParameter3) SetCountry(value string) {
	b.Country = (*CountryCode)(&value)
}

func (b *BreakdownByUserDefinedParameter3) SetCurrency(value string) {
	b.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (b *BreakdownByUserDefinedParameter3) AddUserDefined() *DataFormat2Choice {
	b.UserDefined = new(DataFormat2Choice)
	return b.UserDefined
}

func (b *BreakdownByUserDefinedParameter3) AddCashInForecast() *CashInForecast5 {
	newValue := new(CashInForecast5)
	b.CashInForecast = append(b.CashInForecast, newValue)
	return newValue
}

func (b *BreakdownByUserDefinedParameter3) AddCashOutForecast() *CashOutForecast5 {
	newValue := new(CashOutForecast5)
	b.CashOutForecast = append(b.CashOutForecast, newValue)
	return newValue
}

func (b *BreakdownByUserDefinedParameter3) AddNetCashForecast() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	b.NetCashForecast = append(b.NetCashForecast, newValue)
	return newValue
}
