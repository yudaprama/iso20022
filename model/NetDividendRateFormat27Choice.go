package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat27Choice struct {

	// Number of monetary units specified in a currency.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus2 `xml:"AmtAndRateSts"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus31 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (n *NetDividendRateFormat27Choice) SetAmount(value, currency string) {
	n.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat27Choice) AddAmountAndRateStatus() *AmountAndRateStatus2 {
	n.AmountAndRateStatus = new(AmountAndRateStatus2)
	return n.AmountAndRateStatus
}

func (n *NetDividendRateFormat27Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus31 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus31)
	return n.RateTypeAndAmountAndRateStatus
}

func (n *NetDividendRateFormat27Choice) SetNotSpecifiedRate(value string) {
	n.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
