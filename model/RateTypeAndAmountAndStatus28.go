package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus28 struct {

	// Value expressed as a rate type.
	RateType *RateType38Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus28) AddRateType() *RateType38Choice {
	r.RateType = new(RateType38Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus28) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus28) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}
