package model

// Specifies the cash-in and cash-out flows by currency.
type BreakdownByCurrency1 struct {

	// Currency for which the cash flow is being reported.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy"`

	// Cash movement out of the fund as a result of investment funds transactions, eg, redemptions or switch-out.
	CashOutForecast []*CashOutForecast3 `xml:"CshOutFcst,omitempty"`

	// Cash movement into the fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	CashInForecast []*CashInForecast3 `xml:"CshInFcst,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows specified for the currency.
	NetCashForecast []*NetCashForecast2 `xml:"NetCshFcst,omitempty"`
}

func (b *BreakdownByCurrency1) SetCurrency(value string) {
	b.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (b *BreakdownByCurrency1) AddCashOutForecast() *CashOutForecast3 {
	newValue := new(CashOutForecast3)
	b.CashOutForecast = append(b.CashOutForecast, newValue)
	return newValue
}

func (b *BreakdownByCurrency1) AddCashInForecast() *CashInForecast3 {
	newValue := new(CashInForecast3)
	b.CashInForecast = append(b.CashInForecast, newValue)
	return newValue
}

func (b *BreakdownByCurrency1) AddNetCashForecast() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	b.NetCashForecast = append(b.NetCashForecast, newValue)
	return newValue
}
