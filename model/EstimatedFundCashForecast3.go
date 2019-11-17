package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast3 struct {

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

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Estimated cash movements into the fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	EstimatedCashInForecastDetails []*CashInForecast4 `xml:"EstmtdCshInFcstDtls,omitempty"`

	// Estimated cash movements out of the fund as a result of investment funds transactions, eg, redemptions or switch-out.
	EstimatedCashOutForecastDetails []*CashOutForecast4 `xml:"EstmtdCshOutFcstDtls,omitempty"`

	// Net cash movements to a fund as a result of investment funds transactions.
	EstimatedNetCashForecastDetails []*NetCashForecast2 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast3) SetIdentification(value string) {
	e.Identification = (*Max35Text)(&value)
}

func (e *EstimatedFundCashForecast3) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast3) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast3) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	e.FinancialInstrumentDetails = new(FinancialInstrument9)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast3) SetEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast3) SetPreviousTotalNAV(value, currency string) {
	e.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast3) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast3) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousTotalUnitsNumber
}

func (e *EstimatedFundCashForecast3) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast3) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast3) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast3) AddEstimatedCashInForecastDetails() *CashInForecast4 {
	newValue := new(CashInForecast4)
	e.EstimatedCashInForecastDetails = append(e.EstimatedCashInForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast3) AddEstimatedCashOutForecastDetails() *CashOutForecast4 {
	newValue := new(CashOutForecast4)
	e.EstimatedCashOutForecastDetails = append(e.EstimatedCashOutForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast3) AddEstimatedNetCashForecastDetails() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}
