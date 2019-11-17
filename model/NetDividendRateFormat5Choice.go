package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat5Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus4 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (n *NetDividendRateFormat5Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat5Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus4 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus4)
	return n.RateTypeAndAmountAndRateStatus
}

func (n *NetDividendRateFormat5Choice) SetNotSpecifiedRate(value string) {
	n.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
