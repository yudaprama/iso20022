package model

// An investor's instruction to either subscribe or redeem an amount of money or its equivalent, eg, other assets, into or out of an investment fund.
type InvestmentFundsOrderBreakdown1 struct {

	// Type of order breakdown.
	OrderBreakdownType *FundOrderType5Code `xml:"OrdrBrkdwnTp"`

	// Type of order breakdown.
	ExtendedOrderBreakdownType *Extended350Code `xml:"XtndedOrdrBrkdwnTp"`

	// Portion of the net amount that is attributed to an order type.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (i *InvestmentFundsOrderBreakdown1) SetOrderBreakdownType(value string) {
	i.OrderBreakdownType = (*FundOrderType5Code)(&value)
}

func (i *InvestmentFundsOrderBreakdown1) SetExtendedOrderBreakdownType(value string) {
	i.ExtendedOrderBreakdownType = (*Extended350Code)(&value)
}

func (i *InvestmentFundsOrderBreakdown1) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}
