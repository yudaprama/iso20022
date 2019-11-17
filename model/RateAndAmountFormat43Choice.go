package model

// Choice of format between rate and amount.
type RateAndAmountFormat43Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value is expressed as a currency and amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (r *RateAndAmountFormat43Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat43Choice) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
