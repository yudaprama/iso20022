package model

// Cash dividend amount per equity before deductions or allowances have been made.
type GrossDividendRate2 struct {

	// Type of underlying securities to which the rate is related, eg, underlying security for which an interest is paid.
	RateType *GrossDividendRateType1FormatChoice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (g *GrossDividendRate2) AddRateType() *GrossDividendRateType1FormatChoice {
	g.RateType = new(GrossDividendRateType1FormatChoice)
	return g.RateType
}

func (g *GrossDividendRate2) SetAmount(value, currency string) {
	g.Amount = NewActiveCurrencyAndAmount(value, currency)
}
