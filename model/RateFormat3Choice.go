package model

// Choice between a rate or an unspecified rate.
type RateFormat3Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType5Code `xml:"NotSpcfdRate"`
}

func (r *RateFormat3Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateFormat3Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateType5Code)(&value)
}
