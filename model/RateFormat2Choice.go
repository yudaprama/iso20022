package model

// Choice between a rate or an unspecified rate.
type RateFormat2Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType6Code `xml:"NotSpcfdRate"`
}

func (r *RateFormat2Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateFormat2Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType6Code)(&value)
}
