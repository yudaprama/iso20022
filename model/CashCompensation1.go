package model

// Provides details about the cash compensation such as the fees and the total settlement amount.
type CashCompensation1 struct {

	// Provides the original amount to be settled.
	SettlementAmount *AmountAndDirection20 `xml:"SttlmAmt"`

	// Amount of fees linked to the cash compensation process.
	Fees *AmountAndDirection20 `xml:"Fees,omitempty"`

	// Indicates the value date of the cash compensation.
	ValueDate *ISODate `xml:"ValDt,omitempty"`
}

func (c *CashCompensation1) AddSettlementAmount() *AmountAndDirection20 {
	c.SettlementAmount = new(AmountAndDirection20)
	return c.SettlementAmount
}

func (c *CashCompensation1) AddFees() *AmountAndDirection20 {
	c.Fees = new(AmountAndDirection20)
	return c.Fees
}

func (c *CashCompensation1) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}
