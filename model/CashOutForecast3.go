package model

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast3 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash out amounts, eg,  by transaction or order type.
	CashOutBreakdownDetails []*FundCashOutBreakdown2 `xml:"CshOutBrkdwnDtls,omitempty"`
}

func (c *CashOutForecast3) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast3) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast3) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast3) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashOutForecast3) AddCashOutBreakdownDetails() *FundCashOutBreakdown2 {
	newValue := new(FundCashOutBreakdown2)
	c.CashOutBreakdownDetails = append(c.CashOutBreakdownDetails, newValue)
	return newValue
}
