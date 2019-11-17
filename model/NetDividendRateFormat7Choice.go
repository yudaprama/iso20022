package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat7Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus4 `xml:"RateTpAndAmtAndRateSts"`
}

func (n *NetDividendRateFormat7Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat7Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus4 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus4)
	return n.RateTypeAndAmountAndRateStatus
}
