package model

// Specifies the cash-in and cash-out flows by party.
type BreakdownByParty1 struct {

	// Party, eg, fund management company, for which the cash flow is being reported.
	Party *PartyIdentification2Choice `xml:"Pty"`

	// Additional parameter/s applied to the cash flow by party.
	AdditionalParameters *AdditionalParameters1 `xml:"AddtlParams,omitempty"`

	// Cash movement into the fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	CashInForecast []*CashInForecast3 `xml:"CshInFcst,omitempty"`

	// Cash movement out of the fund as a result of investment funds transactions, eg, redemptions or switch-out.
	CashOutForecast []*CashOutForecast3 `xml:"CshOutFcst,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows specified for the party.
	NetCashForecast []*NetCashForecast2 `xml:"NetCshFcst,omitempty"`
}

func (b *BreakdownByParty1) AddParty() *PartyIdentification2Choice {
	b.Party = new(PartyIdentification2Choice)
	return b.Party
}

func (b *BreakdownByParty1) AddAdditionalParameters() *AdditionalParameters1 {
	b.AdditionalParameters = new(AdditionalParameters1)
	return b.AdditionalParameters
}

func (b *BreakdownByParty1) AddCashInForecast() *CashInForecast3 {
	newValue := new(CashInForecast3)
	b.CashInForecast = append(b.CashInForecast, newValue)
	return newValue
}

func (b *BreakdownByParty1) AddCashOutForecast() *CashOutForecast3 {
	newValue := new(CashOutForecast3)
	b.CashOutForecast = append(b.CashOutForecast, newValue)
	return newValue
}

func (b *BreakdownByParty1) AddNetCashForecast() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	b.NetCashForecast = append(b.NetCashForecast, newValue)
	return newValue
}
