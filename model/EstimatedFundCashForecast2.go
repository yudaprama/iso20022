package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast2 struct {

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm"`

	// Investment fund class to which a cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument5 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous estimated value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousEstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsEstmtdTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous estimated value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousEstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsEstmtdTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Information related to the estimated cash movements reported by pre-defined or user defined criteria.
	SortingCriteriaDetails []*CashSortingCriterion1 `xml:"SrtgCritDtls"`

	// Net cash movements per financial instrument.
	EstimatedNetCashForecastDetails []*NetCashForecast1 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast2) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast2) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast2) AddFinancialInstrumentDetails() *FinancialInstrument5 {
	e.FinancialInstrumentDetails = new(FinancialInstrument5)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast2) SetEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast2) SetPreviousEstimatedTotalNAV(value, currency string) {
	e.PreviousEstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast2) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast2) AddPreviousEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousEstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousEstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast2) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast2) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast2) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast2) AddSortingCriteriaDetails() *CashSortingCriterion1 {
	newValue := new(CashSortingCriterion1)
	e.SortingCriteriaDetails = append(e.SortingCriteriaDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast2) AddEstimatedNetCashForecastDetails() *NetCashForecast1 {
	newValue := new(NetCashForecast1)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}
