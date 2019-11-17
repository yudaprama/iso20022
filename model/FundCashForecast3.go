package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast3 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which a cash flow is related.
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

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Cash movements into the fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	CashInForecastDetails []*CashInForecast4 `xml:"CshInFcstDtls,omitempty"`

	// Cash movements out of the fund as a result of investment funds transactions, eg, redemptions or switch-out.
	CashOutForecastDetails []*CashOutForecast4 `xml:"CshOutFcstDtls,omitempty"`

	// Cash movements from or to a fund as a result of investment funds transactions.
	NetCashForecastDetails []*NetCashForecast2 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *FundCashForecast3) SetIdentification(value string) {
	f.Identification = (*Max35Text)(&value)
}

func (f *FundCashForecast3) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast3) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast3) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	f.FinancialInstrumentDetails = new(FinancialInstrument9)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast3) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast3) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast3) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast3) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast3) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast3) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast3) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast3) AddCashInForecastDetails() *CashInForecast4 {
	newValue := new(CashInForecast4)
	f.CashInForecastDetails = append(f.CashInForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast3) AddCashOutForecastDetails() *CashOutForecast4 {
	newValue := new(CashOutForecast4)
	f.CashOutForecastDetails = append(f.CashOutForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast3) AddNetCashForecastDetails() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}
