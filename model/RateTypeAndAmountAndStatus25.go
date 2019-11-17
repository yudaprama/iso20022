package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus25 struct {

	// Value expressed as a rate type.
	RateType *RateType41Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus25) AddRateType() *RateType41Choice {
	r.RateType = new(RateType41Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus25) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus25) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}
