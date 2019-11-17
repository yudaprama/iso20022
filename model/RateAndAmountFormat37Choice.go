package model

// Choice of format between a rate, an amount or a unspecified rate.
type RateAndAmountFormat37Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Value is expressed as a currency and amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (r *RateAndAmountFormat37Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat37Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RateAndAmountFormat37Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
