package model

// An investor's instruction to either subscribe or redeem an amount of money or its equivalent, for example, other assets, into or out of an investment fund.
type InvestmentFundsOrderBreakdown2 struct {

	// Type of order breakdown.
	OrderBreakdownType *OrderBreakdownType1Choice `xml:"OrdrBrkdwnTp"`

	// Portion of the net amount that is attributed to an order type.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (i *InvestmentFundsOrderBreakdown2) AddOrderBreakdownType() *OrderBreakdownType1Choice {
	i.OrderBreakdownType = new(OrderBreakdownType1Choice)
	return i.OrderBreakdownType
}

func (i *InvestmentFundsOrderBreakdown2) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}
