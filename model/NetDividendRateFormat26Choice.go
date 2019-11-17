package model

// Choice between different formats to express a net dividend.
type NetDividendRateFormat26Choice struct {

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus2 `xml:"AmtAndRateSts"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus36 `xml:"RateTpAndAmtAndRateSts"`
}

func (n *NetDividendRateFormat26Choice) SetAmount(value, currency string) {
	n.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat26Choice) AddAmountAndRateStatus() *AmountAndRateStatus2 {
	n.AmountAndRateStatus = new(AmountAndRateStatus2)
	return n.AmountAndRateStatus
}

func (n *NetDividendRateFormat26Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus36 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus36)
	return n.RateTypeAndAmountAndRateStatus
}
