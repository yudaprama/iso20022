package model

// Choice between an amount or an unspecified rate.
type GrossDividendRateFormat9Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus15 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType13Code `xml:"NotSpcfdRate"`
}

func (g *GrossDividendRateFormat9Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat9Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus15 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus15)
	return g.RateTypeAndAmountAndRateStatus
}

func (g *GrossDividendRateFormat9Choice) SetNotSpecifiedRate(value string) {
	g.NotSpecifiedRate = (*RateType13Code)(&value)
}
