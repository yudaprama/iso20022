package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat22Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus1 `xml:"AmtAndRateSts"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus29 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (n *NetDividendRateFormat22Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat22Choice) AddAmountAndRateStatus() *AmountAndRateStatus1 {
	n.AmountAndRateStatus = new(AmountAndRateStatus1)
	return n.AmountAndRateStatus
}

func (n *NetDividendRateFormat22Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus29 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus29)
	return n.RateTypeAndAmountAndRateStatus
}

func (n *NetDividendRateFormat22Choice) SetNotSpecifiedRate(value string) {
	n.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
