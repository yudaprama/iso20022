package model

// Choice of format between a rate or a rate type and rate or an amount or a unspecified rate.
type RateAndAmountFormat45Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value is expressed as a currency and amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value is expressed as a rate type and a percentage rate.
	RateTypeAndRate *RateTypeAndPercentageRate9 `xml:"RateTpAndRate"`
}

func (r *RateAndAmountFormat45Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat45Choice) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat45Choice) AddRateTypeAndRate() *RateTypeAndPercentageRate9 {
	r.RateTypeAndRate = new(RateTypeAndPercentageRate9)
	return r.RateTypeAndRate
}
