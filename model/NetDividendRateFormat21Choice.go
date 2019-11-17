package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat21Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus1 `xml:"AmtAndRateSts"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus25 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (n *NetDividendRateFormat21Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat21Choice) AddAmountAndRateStatus() *AmountAndRateStatus1 {
	n.AmountAndRateStatus = new(AmountAndRateStatus1)
	return n.AmountAndRateStatus
}

func (n *NetDividendRateFormat21Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus25 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus25)
	return n.RateTypeAndAmountAndRateStatus
}

func (n *NetDividendRateFormat21Choice) SetNotSpecifiedRate(value string) {
	n.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
