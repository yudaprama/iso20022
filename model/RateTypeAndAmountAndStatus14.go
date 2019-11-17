package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus14 struct {

	// Value expressed as a rate type.
	RateType *RateType21Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus14) AddRateType() *RateType21Choice {
	r.RateType = new(RateType21Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus14) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus14) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}
