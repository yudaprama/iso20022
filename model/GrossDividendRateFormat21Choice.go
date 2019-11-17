package model

// Choice of format to express a gross dividend.
type GrossDividendRateFormat21Choice struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus1 `xml:"AmtAndRateSts"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus22 `xml:"RateTpAndAmtAndRateSts"`
}

func (g *GrossDividendRateFormat21Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat21Choice) AddAmountAndRateStatus() *AmountAndRateStatus1 {
	g.AmountAndRateStatus = new(AmountAndRateStatus1)
	return g.AmountAndRateStatus
}

func (g *GrossDividendRateFormat21Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus22 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus22)
	return g.RateTypeAndAmountAndRateStatus
}
