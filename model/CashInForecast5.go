package model

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast5 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash in amounts by transaction and order type.
	CashInBreakdownDetails []*FundCashInBreakdown3 `xml:"CshInBrkdwnDtls,omitempty"`

	// Additional balances for cash amounts and number of units.
	// In an estimated report, the total cash derived from orders placed as a number of units is an estimated cash amount and the total number of units derived from orders placed as a cash amount is an estimated number of units.
	AdditionalBalance *FundBalance1 `xml:"AddtlBal,omitempty"`
}

func (c *CashInForecast5) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast5) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast5) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast5) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashInForecast5) AddCashInBreakdownDetails() *FundCashInBreakdown3 {
	newValue := new(FundCashInBreakdown3)
	c.CashInBreakdownDetails = append(c.CashInBreakdownDetails, newValue)
	return newValue
}

func (c *CashInForecast5) AddAdditionalBalance() *FundBalance1 {
	c.AdditionalBalance = new(FundBalance1)
	return c.AdditionalBalance
}
