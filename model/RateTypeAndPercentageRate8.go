package model

// Specifies the value expressed as a rate type and a percentage rate.
type RateTypeAndPercentageRate8 struct {

	// Value expressed as a rate type.
	RateType *RateType42Choice `xml:"RateTp"`

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (r *RateTypeAndPercentageRate8) AddRateType() *RateType42Choice {
	r.RateType = new(RateType42Choice)
	return r.RateType
}

func (r *RateTypeAndPercentageRate8) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}
