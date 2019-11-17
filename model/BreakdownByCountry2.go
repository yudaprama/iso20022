package model

// Specifies the cash-in and cash-out flows by country.
type BreakdownByCountry2 struct {

	// Country for which the cash flow is being reported.
	Country *CountryCode `xml:"Ctry"`

	// Cash movement into the fund as a result of transactions in shares in an investment fund, for example, subscriptions or switch-ins.
	CashInForecast []*CashInForecast5 `xml:"CshInFcst,omitempty"`

	// Cash movement out of the fund as a result of transactions in shares in an investment fund, for example, redemptions or switch-outs.
	CashOutForecast []*CashOutForecast5 `xml:"CshOutFcst,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows specified for the country.
	NetCashForecast []*NetCashForecast4 `xml:"NetCshFcst,omitempty"`
}

func (b *BreakdownByCountry2) SetCountry(value string) {
	b.Country = (*CountryCode)(&value)
}

func (b *BreakdownByCountry2) AddCashInForecast() *CashInForecast5 {
	newValue := new(CashInForecast5)
	b.CashInForecast = append(b.CashInForecast, newValue)
	return newValue
}

func (b *BreakdownByCountry2) AddCashOutForecast() *CashOutForecast5 {
	newValue := new(CashOutForecast5)
	b.CashOutForecast = append(b.CashOutForecast, newValue)
	return newValue
}

func (b *BreakdownByCountry2) AddNetCashForecast() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	b.NetCashForecast = append(b.NetCashForecast, newValue)
	return newValue
}
