package model

// Choice between different formats to express a net dividend.
type NetDividendRateFormat24Choice struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus1 `xml:"AmtAndRateSts"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus29 `xml:"RateTpAndAmtAndRateSts"`
}

func (n *NetDividendRateFormat24Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat24Choice) AddAmountAndRateStatus() *AmountAndRateStatus1 {
	n.AmountAndRateStatus = new(AmountAndRateStatus1)
	return n.AmountAndRateStatus
}

func (n *NetDividendRateFormat24Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus29 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus29)
	return n.RateTypeAndAmountAndRateStatus
}
