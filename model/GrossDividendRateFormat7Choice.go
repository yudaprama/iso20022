package model

// Choice between an amount or an unspecified rate.
type GrossDividendRateFormat7Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus13 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType13Code `xml:"NotSpcfdRate"`
}

func (g *GrossDividendRateFormat7Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat7Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus13 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus13)
	return g.RateTypeAndAmountAndRateStatus
}

func (g *GrossDividendRateFormat7Choice) SetNotSpecifiedRate(value string) {
	g.NotSpecifiedRate = (*RateType13Code)(&value)
}
