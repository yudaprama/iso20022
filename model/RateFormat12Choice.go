package model

// Choice between a rate or an unspecified rate.
type RateFormat12Choice struct {

	// Value is expressed as a rate.
	Rate *BaseOne14Rate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType5Code `xml:"NotSpcfdRate"`
}

func (r *RateFormat12Choice) SetRate(value string) {
	r.Rate = (*BaseOne14Rate)(&value)
}

func (r *RateFormat12Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateType5Code)(&value)
}
