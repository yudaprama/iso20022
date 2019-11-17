package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast4 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Estimated cash flow by party.
	BreakdownByParty []*BreakdownByParty1 `xml:"BrkdwnByPty,omitempty"`

	// Estimated cash flow by country.
	BreakdownByCountry []*BreakdownByCountry1 `xml:"BrkdwnByCtry,omitempty"`

	// Estimated cash flow by currency.
	BreakdownByCurrency []*BreakdownByCurrency1 `xml:"BrkdwnByCcy,omitempty"`

	// Estimated cash flow by a user defined parameter/s.
	BreakdownByUserDefinedParameter []*BreakdownByUserDefinedParameter1 `xml:"BrkdwnByUsrDfndParam,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Estimated net cash movements per financial instrument.
	EstimatedNetCashForecastDetails []*NetCashForecast2 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast4) SetIdentification(value string) {
	e.Identification = (*Max35Text)(&value)
}

func (e *EstimatedFundCashForecast4) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast4) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast4) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	e.FinancialInstrumentDetails = new(FinancialInstrument9)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast4) SetEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast4) SetPreviousTotalNAV(value, currency string) {
	e.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast4) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast4) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousTotalUnitsNumber
}

func (e *EstimatedFundCashForecast4) AddBreakdownByParty() *BreakdownByParty1 {
	newValue := new(BreakdownByParty1)
	e.BreakdownByParty = append(e.BreakdownByParty, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast4) AddBreakdownByCountry() *BreakdownByCountry1 {
	newValue := new(BreakdownByCountry1)
	e.BreakdownByCountry = append(e.BreakdownByCountry, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast4) AddBreakdownByCurrency() *BreakdownByCurrency1 {
	newValue := new(BreakdownByCurrency1)
	e.BreakdownByCurrency = append(e.BreakdownByCurrency, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast4) AddBreakdownByUserDefinedParameter() *BreakdownByUserDefinedParameter1 {
	newValue := new(BreakdownByUserDefinedParameter1)
	e.BreakdownByUserDefinedParameter = append(e.BreakdownByUserDefinedParameter, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast4) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast4) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast4) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast4) AddEstimatedNetCashForecastDetails() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}
