package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat12Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus16 `xml:"RateTpAndAmtAndRateSts"`
}

func (n *NetDividendRateFormat12Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat12Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus16 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus16)
	return n.RateTypeAndAmountAndRateStatus
}
