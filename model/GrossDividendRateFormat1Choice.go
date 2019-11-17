package model

// Choice between an amount or an unspecified rate.
type GrossDividendRateFormat1Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the gross dividend rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus1 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType12Code `xml:"NotSpcfdRate"`
}

func (g *GrossDividendRateFormat1Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (g *GrossDividendRateFormat1Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus1 {
	g.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus1)
	return g.RateTypeAndAmountAndRateStatus
}

func (g *GrossDividendRateFormat1Choice) SetNotSpecifiedRate(value string) {
	g.NotSpecifiedRate = (*RateType12Code)(&value)
}
