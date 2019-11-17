package model

// Specifies the value expressed as a rate type and a percentage rate.
type RateTypeAndPercentageRate1 struct {

	// Value expressed as a rate type.
	RateType *RateType28Choice `xml:"RateTp"`

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (r *RateTypeAndPercentageRate1) AddRateType() *RateType28Choice {
	r.RateType = new(RateType28Choice)
	return r.RateType
}

func (r *RateTypeAndPercentageRate1) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}
