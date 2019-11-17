package model

// Choice of format to express a gross dividend.
type GrossDividendRateFormat23Choice struct {

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus2 `xml:"AmtAndRateSts"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus30 `xml:"RateTpAndAmtAndRateSts"`
}

func (g *GrossDividendRateFormat23Choice) SetAmount(value, currency string) {
	g.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat23Choice) AddAmountAndRateStatus() *AmountAndRateStatus2 {
	g.AmountAndRateStatus = new(AmountAndRateStatus2)
	return g.AmountAndRateStatus
}

func (g *GrossDividendRateFormat23Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus30 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus30)
	return g.RateTypeAndAmountAndRateStatus
}
