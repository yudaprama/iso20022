package model

// Choice of rate formats.
type AmountAndRateFormat2Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// The rate is not specified.
	NotSpecifiedRate *RateType12FormatChoice `xml:"NotSpcfdRate"`
}

func (a *AmountAndRateFormat2Choice) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndRateFormat2Choice) AddNotSpecifiedRate() *RateType12FormatChoice {
	a.NotSpecifiedRate = new(RateType12FormatChoice)
	return a.NotSpecifiedRate
}
