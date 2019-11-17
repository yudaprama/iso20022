package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus16 struct {

	// Value expressed as a rate type.
	RateType *RateType23Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus16) AddRateType() *RateType23Choice {
	r.RateType = new(RateType23Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus16) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus16) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}
