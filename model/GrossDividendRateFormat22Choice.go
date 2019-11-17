package model

// Choice of format to express a gross dividend.
type GrossDividendRateFormat22Choice struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus1 `xml:"AmtAndRateSts"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus28 `xml:"RateTpAndAmtAndRateSts"`
}

func (g *GrossDividendRateFormat22Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat22Choice) AddAmountAndRateStatus() *AmountAndRateStatus1 {
	g.AmountAndRateStatus = new(AmountAndRateStatus1)
	return g.AmountAndRateStatus
}

func (g *GrossDividendRateFormat22Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus28 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus28)
	return g.RateTypeAndAmountAndRateStatus
}
