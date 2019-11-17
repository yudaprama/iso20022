package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast1 struct {

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

	// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	EstimatedCashInForecastDetails []*CashInForecast2 `xml:"EstmtdCshInFcstDtls,omitempty"`

	// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
	EstimatedCashOutForecastDetails []*CashOutForecast2 `xml:"EstmtdCshOutFcstDtls,omitempty"`

	// Net cash movements to a fund as a result of investment funds transactions.
	EstimatedNetCashForecastDetails []*NetCashForecast1 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast1) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast1) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast1) AddFinancialInstrumentDetails() *FinancialInstrument5 {
	e.FinancialInstrumentDetails = new(FinancialInstrument5)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast1) SetEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast1) SetPreviousEstimatedTotalNAV(value, currency string) {
	e.PreviousEstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast1) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast1) AddPreviousEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousEstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousEstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast1) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast1) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast1) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast1) AddEstimatedCashInForecastDetails() *CashInForecast2 {
	newValue := new(CashInForecast2)
	e.EstimatedCashInForecastDetails = append(e.EstimatedCashInForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast1) AddEstimatedCashOutForecastDetails() *CashOutForecast2 {
	newValue := new(CashOutForecast2)
	e.EstimatedCashOutForecastDetails = append(e.EstimatedCashOutForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast1) AddEstimatedNetCashForecastDetails() *NetCashForecast1 {
	newValue := new(NetCashForecast1)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}
