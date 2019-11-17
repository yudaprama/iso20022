package model

// Choice between a rate or an unspecified rate.
type RateFormat7Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType10Code `xml:"NotSpcfdRate"`
}

func (r *RateFormat7Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateFormat7Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateType10Code)(&value)
}
