package model

// Choice of format between a rate, an amount, index points or a unspecified rate.
type RateAndAmountFormat38Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Value is expressed as a currency and amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Price expressed in index points.
	//
	IndexPoints *DecimalNumber `xml:"IndxPts"`
}

func (r *RateAndAmountFormat38Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat38Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RateAndAmountFormat38Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat38Choice) SetIndexPoints(value string) {
	r.IndexPoints = (*DecimalNumber)(&value)
}
