package model

// Expected trade date and expected settlement date of the order execution.
type ExpectedExecutionDetails1 struct {

	// Expected date or expected date and time at which a price will be applied according to the terms of the prospectus.
	ExpectedTradeDateTime *DateAndDateTimeChoice `xml:"XpctdTradDtTm,omitempty"`

	// Expected date at which the financial instruments will be exchanged against cash.
	ExpectedSettlementDate *ISODate `xml:"XpctdSttlmDt,omitempty"`
}

func (e *ExpectedExecutionDetails1) AddExpectedTradeDateTime() *DateAndDateTimeChoice {
	e.ExpectedTradeDateTime = new(DateAndDateTimeChoice)
	return e.ExpectedTradeDateTime
}

func (e *ExpectedExecutionDetails1) SetExpectedSettlementDate(value string) {
	e.ExpectedSettlementDate = (*ISODate)(&value)
}
