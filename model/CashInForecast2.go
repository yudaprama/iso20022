package model

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast2 struct {

	// Date on which cash is available.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`
}

func (c *CashInForecast2) SetSettlementDate(value string) {
	c.SettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast2) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast2) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast2) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}
