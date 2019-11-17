package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus11 struct {

	// Value expressed as a rate type.
	RateType *RateType17Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus11) AddRateType() *RateType17Choice {
	r.RateType = new(RateType17Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus11) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus11) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}
