package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus26 struct {

	// Value expressed as a rate type.
	RateType *RateType36Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus26) AddRateType() *RateType36Choice {
	r.RateType = new(RateType36Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus26) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus26) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}
