package model

// Choice between a percentage rate or a rate name.
type RateOrName2Choice struct {

	// Pricing expressed as a rate.
	Rate *Rate2 `xml:"Rate"`

	// Pricing expressed as a rate name.
	RateName *RateName2 `xml:"RateNm"`
}

func (r *RateOrName2Choice) AddRate() *Rate2 {
	r.Rate = new(Rate2)
	return r.Rate
}

func (r *RateOrName2Choice) AddRateName() *RateName2 {
	r.RateName = new(RateName2)
	return r.RateName
}
