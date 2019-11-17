package model

// Choice of format to expressed a rate.
type RateFormat1Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value is not specified.
	NotSpecifiedRate *RateType12FormatChoice `xml:"NotSpcfdRate"`
}

func (r *RateFormat1Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateFormat1Choice) AddNotSpecifiedRate() *RateType12FormatChoice {
	r.NotSpecifiedRate = new(RateType12FormatChoice)
	return r.NotSpecifiedRate
}
