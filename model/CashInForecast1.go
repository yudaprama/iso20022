package model

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast1 struct {

	// Date on which cash is available.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash in amounts, eg, by transaction and order type.
	CashInBreakdownDetails []*FundCashInBreakdown1 `xml:"CshInBrkdwnDtls,omitempty"`
}

func (c *CashInForecast1) SetSettlementDate(value string) {
	c.SettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast1) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast1) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast1) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashInForecast1) AddCashInBreakdownDetails() *FundCashInBreakdown1 {
	newValue := new(FundCashInBreakdown1)
	c.CashInBreakdownDetails = append(c.CashInBreakdownDetails, newValue)
	return newValue
}
