package model

// Choice between a rate or a rate type and rate or an unspecified rate.
type RateFormat10Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate is not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Value is expressed as a rate type and a percentage rate.
	RateTypeAndRate *RateTypeAndPercentageRate1 `xml:"RateTpAndRate"`
}

func (r *RateFormat10Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateFormat10Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RateFormat10Choice) AddRateTypeAndRate() *RateTypeAndPercentageRate1 {
	r.RateTypeAndRate = new(RateTypeAndPercentageRate1)
	return r.RateTypeAndRate
}
