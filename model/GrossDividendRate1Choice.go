package model

// Choice of format to express a gross dividend.
type GrossDividendRate1Choice struct {

	// The value of the rate is not specified, eg, the rate is unknown.
	NotSpecifiedRate *RateValueType2FormatChoice `xml:"NotSpcfdRate"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Value is expressed as an amount related to an underlying securities, eg, underlying security for which an interest is paid.
	RateTypeAmount *GrossDividendRate2 `xml:"RateTpAmt"`
}

func (g *GrossDividendRate1Choice) AddNotSpecifiedRate() *RateValueType2FormatChoice {
	g.NotSpecifiedRate = new(RateValueType2FormatChoice)
	return g.NotSpecifiedRate
}

func (g *GrossDividendRate1Choice) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GrossDividendRate1Choice) AddRateTypeAmount() *GrossDividendRate2 {
	g.RateTypeAmount = new(GrossDividendRate2)
	return g.RateTypeAmount
}
