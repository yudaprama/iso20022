package model

// Set of elements qualifying the interest rate.
type Rate2 struct {

	// Indicates the sign of the rate.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`

	// Percentage charged for the use of an amount of money, usually expressed at an annual rate. The interest rate is the ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	Rate *PercentageRate `xml:"Rate"`
}

func (r *Rate2) SetSign(value string) {
	r.Sign = (*PlusOrMinusIndicator)(&value)
}

func (r *Rate2) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}
