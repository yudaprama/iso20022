package model

// Expected trade date and expected settlement date of the order execution.
type ExpectedExecutionDetails4 struct {

	// Expected date or expected date and time at which a price will be applied according to the terms of the prospectus.
	ExpectedTradeDateTime *DateAndDateTimeChoice `xml:"XpctdTradDtTm,omitempty"`

	// Date of a payment, for example, a prepayment date.
	ExpectedCashSettlementDate *ISODate `xml:"XpctdCshSttlmDt,omitempty"`
}

func (e *ExpectedExecutionDetails4) AddExpectedTradeDateTime() *DateAndDateTimeChoice {
	e.ExpectedTradeDateTime = new(DateAndDateTimeChoice)
	return e.ExpectedTradeDateTime
}

func (e *ExpectedExecutionDetails4) SetExpectedCashSettlementDate(value string) {
	e.ExpectedCashSettlementDate = (*ISODate)(&value)
}
