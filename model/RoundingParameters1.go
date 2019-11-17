package model

// Parameters applied to a fractional number.
type RoundingParameters1 struct {

	// Float value specifying the value to which rounding is required, eg, 10 means round to a multiple of 10 units/shares, 0.5 means round to a multiple of 0.5 units/shares.
	RoundingModulus *DecimalNumber `xml:"RndgMdlus,omitempty"`

	// Rounding direction applied to fractional numbers, eg, round up.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn"`
}

func (r *RoundingParameters1) SetRoundingModulus(value string) {
	r.RoundingModulus = (*DecimalNumber)(&value)
}

func (r *RoundingParameters1) SetRoundingDirection(value string) {
	r.RoundingDirection = (*RoundingDirection1Code)(&value)
}
