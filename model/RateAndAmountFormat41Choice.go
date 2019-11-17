package model

// Choice of format between a rate or a rate type and rate or an amount or a unspecified rate.
type RateAndAmountFormat41Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Value is expressed as a currency and amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value is expressed as a rate type and a percentage rate.
	RateTypeAndRate *RateTypeAndPercentageRate8 `xml:"RateTpAndRate"`
}

func (r *RateAndAmountFormat41Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat41Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RateAndAmountFormat41Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat41Choice) AddRateTypeAndRate() *RateTypeAndPercentageRate8 {
	r.RateTypeAndRate = new(RateTypeAndPercentageRate8)
	return r.RateTypeAndRate
}
