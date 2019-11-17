package model

// Choice between a rate or an unspecified rate.
type RateFormat5Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType9Code `xml:"NotSpcfdRate"`
}

func (r *RateFormat5Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateFormat5Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateType9Code)(&value)
}
