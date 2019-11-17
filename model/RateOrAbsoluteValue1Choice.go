package model

// Choice between a rate or an absolute value.
type RateOrAbsoluteValue1Choice struct {

	// A rate expressed as a percentage.
	RateValue *PercentageRate `xml:"RateVal"`

	// Absolute value determined with a number.
	AbsoluteValue *Number `xml:"AbsVal"`
}

func (r *RateOrAbsoluteValue1Choice) SetRateValue(value string) {
	r.RateValue = (*PercentageRate)(&value)
}

func (r *RateOrAbsoluteValue1Choice) SetAbsoluteValue(value string) {
	r.AbsoluteValue = (*Number)(&value)
}
