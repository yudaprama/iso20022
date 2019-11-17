package model

// Expected trade date and expected settlement date of the order execution.
type ExpectedExecutionDetails2 struct {

	// Expected date or expected date and time at which a price will be applied according to the terms of the prospectus.
	ExpectedTradeDateTime *DateAndDateTimeChoice `xml:"XpctdTradDtTm,omitempty"`

	// Expected date at which the financial instruments will be exchanged against cash.
	ExpectedCashSettlementDate *ISODate `xml:"XpctdCshSttlmDt,omitempty"`
}

func (e *ExpectedExecutionDetails2) AddExpectedTradeDateTime() *DateAndDateTimeChoice {
	e.ExpectedTradeDateTime = new(DateAndDateTimeChoice)
	return e.ExpectedTradeDateTime
}

func (e *ExpectedExecutionDetails2) SetExpectedCashSettlementDate(value string) {
	e.ExpectedCashSettlementDate = (*ISODate)(&value)
}
