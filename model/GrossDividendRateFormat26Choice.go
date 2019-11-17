package model

// Choice between an amount or an unspecified rate.
type GrossDividendRateFormat26Choice struct {

	// Number of monetary units specified in a currency.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus2 `xml:"AmtAndRateSts"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus35 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType13Code `xml:"NotSpcfdRate"`
}

func (g *GrossDividendRateFormat26Choice) SetAmount(value, currency string) {
	g.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat26Choice) AddAmountAndRateStatus() *AmountAndRateStatus2 {
	g.AmountAndRateStatus = new(AmountAndRateStatus2)
	return g.AmountAndRateStatus
}

func (g *GrossDividendRateFormat26Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus35 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus35)
	return g.RateTypeAndAmountAndRateStatus
}

func (g *GrossDividendRateFormat26Choice) SetNotSpecifiedRate(value string) {
	g.NotSpecifiedRate = (*RateType13Code)(&value)
}
