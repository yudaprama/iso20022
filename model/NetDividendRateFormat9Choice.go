package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat9Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus14 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (n *NetDividendRateFormat9Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat9Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus14 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus14)
	return n.RateTypeAndAmountAndRateStatus
}

func (n *NetDividendRateFormat9Choice) SetNotSpecifiedRate(value string) {
	n.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
