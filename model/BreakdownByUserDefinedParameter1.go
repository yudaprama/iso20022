package model

// Specifies the cash-in and cash-out flows by a user defined parameter/s.
type BreakdownByUserDefinedParameter1 struct {

	// Party for which the cash flow is being reported.
	Party *PartyIdentification2Choice `xml:"Pty,omitempty"`

	// Country for which the cash flow is being reported.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Currency for which the cash flow is being reported.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Parameter for which the cash flow is being reported.
	UserDefined *DataFormat2Choice `xml:"UsrDfnd,omitempty"`

	// Cash movement into the fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	CashInForecast []*CashInForecast3 `xml:"CshInFcst,omitempty"`

	// Cash movement out of the fund as a result of investment funds transactions, eg, redemptions or switch-out.
	CashOutForecast []*CashOutForecast3 `xml:"CshOutFcst,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows specified for the user defined parameter.
	NetCashForecast []*NetCashForecast2 `xml:"NetCshFcst,omitempty"`
}

func (b *BreakdownByUserDefinedParameter1) AddParty() *PartyIdentification2Choice {
	b.Party = new(PartyIdentification2Choice)
	return b.Party
}

func (b *BreakdownByUserDefinedParameter1) SetCountry(value string) {
	b.Country = (*CountryCode)(&value)
}

func (b *BreakdownByUserDefinedParameter1) SetCurrency(value string) {
	b.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (b *BreakdownByUserDefinedParameter1) AddUserDefined() *DataFormat2Choice {
	b.UserDefined = new(DataFormat2Choice)
	return b.UserDefined
}

func (b *BreakdownByUserDefinedParameter1) AddCashInForecast() *CashInForecast3 {
	newValue := new(CashInForecast3)
	b.CashInForecast = append(b.CashInForecast, newValue)
	return newValue
}

func (b *BreakdownByUserDefinedParameter1) AddCashOutForecast() *CashOutForecast3 {
	newValue := new(CashOutForecast3)
	b.CashOutForecast = append(b.CashOutForecast, newValue)
	return newValue
}

func (b *BreakdownByUserDefinedParameter1) AddNetCashForecast() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	b.NetCashForecast = append(b.NetCashForecast, newValue)
	return newValue
}
