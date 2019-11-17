package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus5 struct {

	// Value expressed as a rate type.
	RateType *RateType10Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus5) AddRateType() *RateType10Choice {
	r.RateType = new(RateType10Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus5) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus5) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}
