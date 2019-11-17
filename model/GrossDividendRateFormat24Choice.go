package model

// Choice of format to express a gross dividend.
type GrossDividendRateFormat24Choice struct {

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus2 `xml:"AmtAndRateSts"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus35 `xml:"RateTpAndAmtAndRateSts"`
}

func (g *GrossDividendRateFormat24Choice) SetAmount(value, currency string) {
	g.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat24Choice) AddAmountAndRateStatus() *AmountAndRateStatus2 {
	g.AmountAndRateStatus = new(AmountAndRateStatus2)
	return g.AmountAndRateStatus
}

func (g *GrossDividendRateFormat24Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus35 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus35)
	return g.RateTypeAndAmountAndRateStatus
}
