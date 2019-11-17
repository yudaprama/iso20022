package model

// Choice between different formats to express a net dividend.
type NetDividendRateFormat2Choice struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus4 `xml:"RateTpAndAmtAndRateSts"`
}

func (n *NetDividendRateFormat2Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat2Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus4 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus4)
	return n.RateTypeAndAmountAndRateStatus
}
