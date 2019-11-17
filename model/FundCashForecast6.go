package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast6 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which the price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of investment fund class units that have been issued.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	TotalNAVChangeRate *PercentageRate `xml:"TtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Information about the designation of the share class currency, that is, whether it is for onshore or offshore purposes and other information that may be required. This is typically only required for CNY funds.
	CurrencyStatus *CurrencyDesignation1 `xml:"CcySts,omitempty"`

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Price per unit of the trade date.
	Price *UnitPrice19 `xml:"Pric,omitempty"`

	// Foreign exchange rate.
	ForeignExchangeRate *ForeignExchangeTerms19 `xml:"FXRate,omitempty"`

	// Net cash flow expressed as a percentage of the total NAV for the share class.
	PercentageOfShareClassTotalNAV *PercentageRate `xml:"PctgOfShrClssTtlNAV,omitempty"`

	// Cash flow by party.
	BreakdownByParty []*BreakdownByParty3 `xml:"BrkdwnByPty,omitempty"`

	// Cash flow by country.
	BreakdownByCountry []*BreakdownByCountry2 `xml:"BrkdwnByCtry,omitempty"`

	// Cash flow by currency.
	BreakdownByCurrency []*BreakdownByCurrency2 `xml:"BrkdwnByCcy,omitempty"`

	// Cash flow by a user defined parameter/s.
	BreakdownByUserDefinedParameter []*BreakdownByUserDefinedParameter3 `xml:"BrkdwnByUsrDfndParam,omitempty"`

	// Net cash movements per financial instrument.
	NetCashForecastDetails []*NetCashForecast4 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *FundCashForecast6) SetIdentification(value string) {
	f.Identification = (*Max35Text)(&value)
}

func (f *FundCashForecast6) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast6) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast6) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	f.FinancialInstrumentDetails = new(FinancialInstrument9)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast6) AddTotalNAV(value, currency string) {
	f.TotalNAV = append(f.TotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (f *FundCashForecast6) AddPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = append(f.PreviousTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (f *FundCashForecast6) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast6) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast6) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast6) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast6) AddCurrencyStatus() *CurrencyDesignation1 {
	f.CurrencyStatus = new(CurrencyDesignation1)
	return f.CurrencyStatus
}

func (f *FundCashForecast6) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast6) AddPrice() *UnitPrice19 {
	f.Price = new(UnitPrice19)
	return f.Price
}

func (f *FundCashForecast6) AddForeignExchangeRate() *ForeignExchangeTerms19 {
	f.ForeignExchangeRate = new(ForeignExchangeTerms19)
	return f.ForeignExchangeRate
}

func (f *FundCashForecast6) SetPercentageOfShareClassTotalNAV(value string) {
	f.PercentageOfShareClassTotalNAV = (*PercentageRate)(&value)
}

func (f *FundCashForecast6) AddBreakdownByParty() *BreakdownByParty3 {
	newValue := new(BreakdownByParty3)
	f.BreakdownByParty = append(f.BreakdownByParty, newValue)
	return newValue
}

func (f *FundCashForecast6) AddBreakdownByCountry() *BreakdownByCountry2 {
	newValue := new(BreakdownByCountry2)
	f.BreakdownByCountry = append(f.BreakdownByCountry, newValue)
	return newValue
}

func (f *FundCashForecast6) AddBreakdownByCurrency() *BreakdownByCurrency2 {
	newValue := new(BreakdownByCurrency2)
	f.BreakdownByCurrency = append(f.BreakdownByCurrency, newValue)
	return newValue
}

func (f *FundCashForecast6) AddBreakdownByUserDefinedParameter() *BreakdownByUserDefinedParameter3 {
	newValue := new(BreakdownByUserDefinedParameter3)
	f.BreakdownByUserDefinedParameter = append(f.BreakdownByUserDefinedParameter, newValue)
	return newValue
}

func (f *FundCashForecast6) AddNetCashForecastDetails() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}
