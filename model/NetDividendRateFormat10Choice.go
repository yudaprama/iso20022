package model

// Choice between different formats to express a net dividend.
type NetDividendRateFormat10Choice struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus14 `xml:"RateTpAndAmtAndRateSts"`
}

func (n *NetDividendRateFormat10Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat10Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus14 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus14)
	return n.RateTypeAndAmountAndRateStatus
}
