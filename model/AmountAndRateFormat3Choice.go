package model

// Choice of rate formats.
type AmountAndRateFormat3Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// The rate is not specified.
	NotSpecifiedRate *RateValueType6FormatChoice `xml:"NotSpcfdRate"`
}

func (a *AmountAndRateFormat3Choice) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndRateFormat3Choice) AddNotSpecifiedRate() *RateValueType6FormatChoice {
	a.NotSpecifiedRate = new(RateValueType6FormatChoice)
	return a.NotSpecifiedRate
}
