package model

// Specifies the cash-in and cash-out flows by currency.
type BreakdownByCurrency2 struct {

	// Currency for which the cash flow is being reported.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy"`

	// Cash movement out of the fund as a result of transactions in shares in an investment fund, for example, redemptions or switch-outs.
	CashOutForecast []*CashOutForecast5 `xml:"CshOutFcst,omitempty"`

	// Cash movement into the fund as a result of transactions in shares in an investment fund, for example, subscriptions or switch-ins.
	CashInForecast []*CashInForecast5 `xml:"CshInFcst,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows specified for the currency.
	NetCashForecast []*NetCashForecast4 `xml:"NetCshFcst,omitempty"`
}

func (b *BreakdownByCurrency2) SetCurrency(value string) {
	b.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (b *BreakdownByCurrency2) AddCashOutForecast() *CashOutForecast5 {
	newValue := new(CashOutForecast5)
	b.CashOutForecast = append(b.CashOutForecast, newValue)
	return newValue
}

func (b *BreakdownByCurrency2) AddCashInForecast() *CashInForecast5 {
	newValue := new(CashInForecast5)
	b.CashInForecast = append(b.CashInForecast, newValue)
	return newValue
}

func (b *BreakdownByCurrency2) AddNetCashForecast() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	b.NetCashForecast = append(b.NetCashForecast, newValue)
	return newValue
}
