package model

// Choice of format between a rate or a rate type and rate or an amount or a unspecified rate.
type RateAndAmountFormat21Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value is expressed as a currency and amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value is expressed as a rate type and a percentage rate.
	RateTypeAndRate *RateTypeAndPercentageRate1 `xml:"RateTpAndRate"`
}

func (r *RateAndAmountFormat21Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat21Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat21Choice) AddRateTypeAndRate() *RateTypeAndPercentageRate1 {
	r.RateTypeAndRate = new(RateTypeAndPercentageRate1)
	return r.RateTypeAndRate
}
