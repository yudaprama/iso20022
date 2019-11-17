package model

// Choice of format to express a gross dividend.
type GrossDividendRateFormat10Choice struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus15 `xml:"RateTpAndAmtAndRateSts"`
}

func (g *GrossDividendRateFormat10Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat10Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus15 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus15)
	return g.RateTypeAndAmountAndRateStatus
}
