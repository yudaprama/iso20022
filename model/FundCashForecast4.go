package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast4 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of investment fund class units that have been issued.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	TotalNAVChangeRate *PercentageRate `xml:"TtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Net cash movements per financial instrument.
	NetCashForecastDetails []*NetCashForecast2 `xml:"NetCshFcstDtls,omitempty"`

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Cash flow by country.
	BreakdownByCountry []*BreakdownByCountry1 `xml:"BrkdwnByCtry,omitempty"`

	// Cash flow by currency.
	BreakdownByCurrency []*BreakdownByCurrency1 `xml:"BrkdwnByCcy,omitempty"`

	// Cash flow by party.
	BreakdownByParty []*BreakdownByParty1 `xml:"BrkdwnByPty,omitempty"`

	// Cash flow by a user defined parameter/s.
	BreakdownByUserDefinedParameter []*BreakdownByUserDefinedParameter1 `xml:"BrkdwnByUsrDfndParam,omitempty"`
}

func (f *FundCashForecast4) SetIdentification(value string) {
	f.Identification = (*Max35Text)(&value)
}

func (f *FundCashForecast4) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast4) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast4) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	f.FinancialInstrumentDetails = new(FinancialInstrument9)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast4) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast4) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast4) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast4) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast4) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast4) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast4) AddNetCashForecastDetails() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast4) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast4) AddBreakdownByCountry() *BreakdownByCountry1 {
	newValue := new(BreakdownByCountry1)
	f.BreakdownByCountry = append(f.BreakdownByCountry, newValue)
	return newValue
}

func (f *FundCashForecast4) AddBreakdownByCurrency() *BreakdownByCurrency1 {
	newValue := new(BreakdownByCurrency1)
	f.BreakdownByCurrency = append(f.BreakdownByCurrency, newValue)
	return newValue
}

func (f *FundCashForecast4) AddBreakdownByParty() *BreakdownByParty1 {
	newValue := new(BreakdownByParty1)
	f.BreakdownByParty = append(f.BreakdownByParty, newValue)
	return newValue
}

func (f *FundCashForecast4) AddBreakdownByUserDefinedParameter() *BreakdownByUserDefinedParameter1 {
	newValue := new(BreakdownByUserDefinedParameter1)
	f.BreakdownByUserDefinedParameter = append(f.BreakdownByUserDefinedParameter, newValue)
	return newValue
}
