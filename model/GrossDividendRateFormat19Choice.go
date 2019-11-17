package model

// Choice between an amount or an unspecified rate.
type GrossDividendRateFormat19Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus1 `xml:"AmtAndRateSts"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus22 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType13Code `xml:"NotSpcfdRate"`
}

func (g *GrossDividendRateFormat19Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat19Choice) AddAmountAndRateStatus() *AmountAndRateStatus1 {
	g.AmountAndRateStatus = new(AmountAndRateStatus1)
	return g.AmountAndRateStatus
}

func (g *GrossDividendRateFormat19Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus22 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus22)
	return g.RateTypeAndAmountAndRateStatus
}

func (g *GrossDividendRateFormat19Choice) SetNotSpecifiedRate(value string) {
	g.NotSpecifiedRate = (*RateType13Code)(&value)
}
