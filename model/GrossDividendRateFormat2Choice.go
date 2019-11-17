package model

// Choice of format to express a gross dividend.
type GrossDividendRateFormat2Choice struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus1 `xml:"RateTpAndAmtAndRateSts"`
}

func (g *GrossDividendRateFormat2Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat2Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus1 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus1)
	return g.RateTypeAndAmountAndRateStatus
}
