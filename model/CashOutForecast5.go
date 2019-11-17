package model

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast5 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash out amounts by transaction and order type.
	CashOutBreakdownDetails []*FundCashOutBreakdown3 `xml:"CshOutBrkdwnDtls,omitempty"`

	// Additional balances for cash amounts and number of units.
	// In an estimated report, the total cash derived from orders placed as a number of units is an estimated cash amount and the total number of units derived from orders placed as a cash amount is an estimated number of units.
	AdditionalBalance *FundBalance1 `xml:"AddtlBal,omitempty"`
}

func (c *CashOutForecast5) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast5) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast5) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast5) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashOutForecast5) AddCashOutBreakdownDetails() *FundCashOutBreakdown3 {
	newValue := new(FundCashOutBreakdown3)
	c.CashOutBreakdownDetails = append(c.CashOutBreakdownDetails, newValue)
	return newValue
}

func (c *CashOutForecast5) AddAdditionalBalance() *FundBalance1 {
	c.AdditionalBalance = new(FundBalance1)
	return c.AdditionalBalance
}
