package model

// Choice of format between a rate or a rate type and rate or an amount or a unspecified rate.
type RateAndAmountFormat47Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Value is expressed as a currency and amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value is expressed as a rate type and a percentage rate.
	RateTypeAndRate *RateTypeAndPercentageRate9 `xml:"RateTpAndRate"`
}

func (r *RateAndAmountFormat47Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat47Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RateAndAmountFormat47Choice) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat47Choice) AddRateTypeAndRate() *RateTypeAndPercentageRate9 {
	r.RateTypeAndRate = new(RateTypeAndPercentageRate9)
	return r.RateTypeAndRate
}
