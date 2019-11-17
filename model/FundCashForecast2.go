package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast2 struct {

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which a cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument5 `xml:"FinInstrmDtls"`

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

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Information related to the cash movements reported by pre-defined or user defined criteria.
	SortingCriteriaDetails []*CashSortingCriterion2 `xml:"SrtgCritDtls"`

	// Net cash movements per financial instrument.
	NetCashForecastDetails []*NetCashForecast1 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *FundCashForecast2) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast2) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast2) AddFinancialInstrumentDetails() *FinancialInstrument5 {
	f.FinancialInstrumentDetails = new(FinancialInstrument5)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast2) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast2) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast2) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast2) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast2) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast2) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast2) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast2) AddSortingCriteriaDetails() *CashSortingCriterion2 {
	newValue := new(CashSortingCriterion2)
	f.SortingCriteriaDetails = append(f.SortingCriteriaDetails, newValue)
	return newValue
}

func (f *FundCashForecast2) AddNetCashForecastDetails() *NetCashForecast1 {
	newValue := new(NetCashForecast1)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}
