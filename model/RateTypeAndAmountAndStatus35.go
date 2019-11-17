package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus35 struct {

	// Value expressed as a rate type.
	RateType *RateType49Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus35) AddRateType() *RateType49Choice {
	r.RateType = new(RateType49Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus35) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus35) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}
