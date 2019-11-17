package model

// Choice of format between a rate, an amount or a unspecified rate.
type RateAndAmountFormat3Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType6Code `xml:"NotSpcfdRate"`

	// Value is expressed as a currency and amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (r *RateAndAmountFormat3Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat3Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType6Code)(&value)
}

func (r *RateAndAmountFormat3Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
