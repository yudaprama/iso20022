package model

// Choice between a percentage rate or a rate name.
type RateOrName1Choice struct {

	// Pricing expressed as a rate.
	Rate *Rate2 `xml:"Rate"`

	// Pricing expressed as a rate name.
	RateName *RateName1 `xml:"RateNm"`
}

func (r *RateOrName1Choice) AddRate() *Rate2 {
	r.Rate = new(Rate2)
	return r.Rate
}

func (r *RateOrName1Choice) AddRateName() *RateName1 {
	r.RateName = new(RateName1)
	return r.RateName
}
