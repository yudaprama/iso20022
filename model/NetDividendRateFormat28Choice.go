package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat28Choice struct {

	// Number of monetary units specified in a currency.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus2 `xml:"AmtAndRateSts"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus36 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (n *NetDividendRateFormat28Choice) SetAmount(value, currency string) {
	n.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat28Choice) AddAmountAndRateStatus() *AmountAndRateStatus2 {
	n.AmountAndRateStatus = new(AmountAndRateStatus2)
	return n.AmountAndRateStatus
}

func (n *NetDividendRateFormat28Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus36 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus36)
	return n.RateTypeAndAmountAndRateStatus
}

func (n *NetDividendRateFormat28Choice) SetNotSpecifiedRate(value string) {
	n.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
