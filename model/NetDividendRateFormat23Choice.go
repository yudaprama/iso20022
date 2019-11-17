package model

// Choice between an amount or an unspecified rate.
type NetDividendRateFormat23Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus1 `xml:"AmtAndRateSts"`

	// Specifies different formats for the net dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus25 `xml:"RateTpAndAmtAndRateSts"`
}

func (n *NetDividendRateFormat23Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (n *NetDividendRateFormat23Choice) AddAmountAndRateStatus() *AmountAndRateStatus1 {
	n.AmountAndRateStatus = new(AmountAndRateStatus1)
	return n.AmountAndRateStatus
}

func (n *NetDividendRateFormat23Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus25 {
	n.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus25)
	return n.RateTypeAndAmountAndRateStatus
}
