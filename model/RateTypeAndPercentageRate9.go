package model

// Specifies the value expressed as a rate type and a percentage rate.
type RateTypeAndPercentageRate9 struct {

	// Value expressed as a rate type.
	RateType *RateType46Choice `xml:"RateTp"`

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (r *RateTypeAndPercentageRate9) AddRateType() *RateType46Choice {
	r.RateType = new(RateType46Choice)
	return r.RateType
}

func (r *RateTypeAndPercentageRate9) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}
