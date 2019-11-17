package model

// Specifies the cash-in and cash-out flows by party.
type BreakdownByParty3 struct {

	// Party, for example, fund management company, for which the cash flow is being reported.
	Party *InvestmentAccount42 `xml:"Pty"`

	// Additional parameter/s applied to the cash flow by party.
	AdditionalParameters *AdditionalParameters1 `xml:"AddtlParams,omitempty"`

	// Cash movement into the fund as a result of transactions in shares in an investment fund, for example, subscriptions or switch-ins.
	CashInForecast []*CashInForecast5 `xml:"CshInFcst,omitempty"`

	// Cash movement out of the fund as a result of transactions in shares in an investment fund, for example, redemptions or switch-outs.
	CashOutForecast []*CashOutForecast5 `xml:"CshOutFcst,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows specified for the party.
	NetCashForecast []*NetCashForecast4 `xml:"NetCshFcst,omitempty"`
}

func (b *BreakdownByParty3) AddParty() *InvestmentAccount42 {
	b.Party = new(InvestmentAccount42)
	return b.Party
}

func (b *BreakdownByParty3) AddAdditionalParameters() *AdditionalParameters1 {
	b.AdditionalParameters = new(AdditionalParameters1)
	return b.AdditionalParameters
}

func (b *BreakdownByParty3) AddCashInForecast() *CashInForecast5 {
	newValue := new(CashInForecast5)
	b.CashInForecast = append(b.CashInForecast, newValue)
	return newValue
}

func (b *BreakdownByParty3) AddCashOutForecast() *CashOutForecast5 {
	newValue := new(CashOutForecast5)
	b.CashOutForecast = append(b.CashOutForecast, newValue)
	return newValue
}

func (b *BreakdownByParty3) AddNetCashForecast() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	b.NetCashForecast = append(b.NetCashForecast, newValue)
	return newValue
}
