package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus32 struct {

	// Value expressed as a rate type.
	RateType *RateType45Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus32) AddRateType() *RateType45Choice {
	r.RateType = new(RateType45Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus32) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus32) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}
