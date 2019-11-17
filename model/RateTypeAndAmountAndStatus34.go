package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus34 struct {

	// Value expressed as a rate type.
	RateType *RateType48Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus34) AddRateType() *RateType48Choice {
	r.RateType = new(RateType48Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus34) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus34) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}
