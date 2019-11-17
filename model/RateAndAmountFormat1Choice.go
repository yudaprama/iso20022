package model

// Choice of format between rate, amount and not specified.
type RateAndAmountFormat1Choice struct {

	// The value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// The value is expressed as a currency and amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// No value is specified.
	NotSpecifiedRate *RateType12FormatChoice `xml:"NotSpcfdRate"`
}

func (r *RateAndAmountFormat1Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat1Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RateAndAmountFormat1Choice) AddNotSpecifiedRate() *RateType12FormatChoice {
	r.NotSpecifiedRate = new(RateType12FormatChoice)
	return r.NotSpecifiedRate
}
