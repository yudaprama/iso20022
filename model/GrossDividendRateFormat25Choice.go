package model

// Choice between an amount or an unspecified rate.
type GrossDividendRateFormat25Choice struct {

	// Number of monetary units specified in a currency.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies an amount and a rate status.
	AmountAndRateStatus *AmountAndRateStatus2 `xml:"AmtAndRateSts"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus30 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType13Code `xml:"NotSpcfdRate"`
}

func (g *GrossDividendRateFormat25Choice) SetAmount(value, currency string) {
	g.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat25Choice) AddAmountAndRateStatus() *AmountAndRateStatus2 {
	g.AmountAndRateStatus = new(AmountAndRateStatus2)
	return g.AmountAndRateStatus
}

func (g *GrossDividendRateFormat25Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus30 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus30)
	return g.RateTypeAndAmountAndRateStatus
}

func (g *GrossDividendRateFormat25Choice) SetNotSpecifiedRate(value string) {
	g.NotSpecifiedRate = (*RateType13Code)(&value)
}
