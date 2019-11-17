package model

// Choice between an amount or an unspecified rate.
type GrossDividendRateFormat20Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus1 `xml:"AmtAndRateSts"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus28 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType13Code `xml:"NotSpcfdRate"`
}

func (g *GrossDividendRateFormat20Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat20Choice) AddAmountAndRateStatus() *AmountAndRateStatus1 {
	g.AmountAndRateStatus = new(AmountAndRateStatus1)
	return g.AmountAndRateStatus
}

func (g *GrossDividendRateFormat20Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus28 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus28)
	return g.RateTypeAndAmountAndRateStatus
}

func (g *GrossDividendRateFormat20Choice) SetNotSpecifiedRate(value string) {
	g.NotSpecifiedRate = (*RateType13Code)(&value)
}
