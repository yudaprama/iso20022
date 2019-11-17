package model

// Choice between a rate or a rate type and rate or an unspecified rate.
type RateFormat11Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value is expressed as a rate type and a percentage rate.
	RateTypeAndRate *RateTypeAndPercentageRate1 `xml:"RateTpAndRate"`
}

func (r *RateFormat11Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateFormat11Choice) AddRateTypeAndRate() *RateTypeAndPercentageRate1 {
	r.RateTypeAndRate = new(RateTypeAndPercentageRate1)
	return r.RateTypeAndRate
}
