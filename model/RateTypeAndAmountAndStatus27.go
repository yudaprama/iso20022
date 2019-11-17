package model

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus27 struct {

	// Value expressed as a rate type.
	RateType *RateType37Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus27) AddRateType() *RateType37Choice {
	r.RateType = new(RateType37Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus27) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus27) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}
