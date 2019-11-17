package model

// Cash movements into or out of a fund as a result of investment funds transactions, for example, subscriptions, redemptions or switches.
type CashInOutForecast7 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Amount of the cash flow.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (c *CashInOutForecast7) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashInOutForecast7) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
