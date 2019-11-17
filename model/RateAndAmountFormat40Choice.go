package model

// Choice of format between a rate or a rate type and rate or an amount or a unspecified rate.
type RateAndAmountFormat40Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value is expressed as a currency and amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value is expressed as a rate type and a percentage rate.
	RateTypeAndRate *RateTypeAndPercentageRate8 `xml:"RateTpAndRate"`
}

func (r *RateAndAmountFormat40Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat40Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat40Choice) AddRateTypeAndRate() *RateTypeAndPercentageRate8 {
	r.RateTypeAndRate = new(RateTypeAndPercentageRate8)
	return r.RateTypeAndRate
}
