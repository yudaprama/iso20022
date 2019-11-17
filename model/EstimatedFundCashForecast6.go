package model

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast6 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price will be applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which the price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Information about the designation of the share class currency, that is, whether it is for onshore or offshore purposes and other information that may be required. This is typically only required for CNY funds.
	CurrencyStatus *CurrencyDesignation1 `xml:"CcySts,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Price per unit of the previous trade date.
	Price *UnitPrice19 `xml:"Pric,omitempty"`

	// Foreign exchange rate.
	ForeignExchangeRate *ForeignExchangeTerms19 `xml:"FXRate,omitempty"`

	// Estimated net cash flow expressed as a percentage of the previous total NAV for the share class.
	EstimatedPercentageOfShareClassTotalNAV *PercentageRate `xml:"EstmtdPctgOfShrClssTtlNAV,omitempty"`

	// Estimated cash movements into the fund as a result of transactions in shares in an investment fund, for example, subscriptions or switch-ins.
	EstimatedCashInForecastDetails []*CashInForecast6 `xml:"EstmtdCshInFcstDtls,omitempty"`

	// Estimated cash movements out of the fund as a result of transactions in shares in an investment fund, for example, redemptions or switch-outs.
	EstimatedCashOutForecastDetails []*CashOutForecast6 `xml:"EstmtdCshOutFcstDtls,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows.
	EstimatedNetCashForecastDetails []*NetCashForecast4 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast6) SetIdentification(value string) {
	e.Identification = (*Max35Text)(&value)
}

func (e *EstimatedFundCashForecast6) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast6) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast6) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	e.FinancialInstrumentDetails = new(FinancialInstrument9)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast6) AddEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = append(e.EstimatedTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (e *EstimatedFundCashForecast6) AddPreviousTotalNAV(value, currency string) {
	e.PreviousTotalNAV = append(e.PreviousTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (e *EstimatedFundCashForecast6) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast6) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousTotalUnitsNumber
}

func (e *EstimatedFundCashForecast6) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast6) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast6) AddCurrencyStatus() *CurrencyDesignation1 {
	e.CurrencyStatus = new(CurrencyDesignation1)
	return e.CurrencyStatus
}

func (e *EstimatedFundCashForecast6) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast6) AddPrice() *UnitPrice19 {
	e.Price = new(UnitPrice19)
	return e.Price
}

func (e *EstimatedFundCashForecast6) AddForeignExchangeRate() *ForeignExchangeTerms19 {
	e.ForeignExchangeRate = new(ForeignExchangeTerms19)
	return e.ForeignExchangeRate
}

func (e *EstimatedFundCashForecast6) SetEstimatedPercentageOfShareClassTotalNAV(value string) {
	e.EstimatedPercentageOfShareClassTotalNAV = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast6) AddEstimatedCashInForecastDetails() *CashInForecast6 {
	newValue := new(CashInForecast6)
	e.EstimatedCashInForecastDetails = append(e.EstimatedCashInForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast6) AddEstimatedCashOutForecastDetails() *CashOutForecast6 {
	newValue := new(CashOutForecast6)
	e.EstimatedCashOutForecastDetails = append(e.EstimatedCashOutForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast6) AddEstimatedNetCashForecastDetails() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}
