package model

// Choice of format to express a gross dividend.
type GrossDividendRateFormat8Choice struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus13 `xml:"RateTpAndAmtAndRateSts"`
}

func (g *GrossDividendRateFormat8Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat8Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus13 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus13)
	return g.RateTypeAndAmountAndRateStatus
}
