package model

// Margin over or under an index that determines the repurchase rate expressed as a rate or an amount.
type SpreadRate1 struct {

	// Specifies the sign of the rate.
	Sign *PlusOrMinusIndicator `xml:"Sgn"`

	// Specifies if the spreadf is expressed as an amount or a rate.
	RateOrAmount *AmountOrRate1Choice `xml:"RateOrAmt"`
}

func (s *SpreadRate1) SetSign(value string) {
	s.Sign = (*PlusOrMinusIndicator)(&value)
}

func (s *SpreadRate1) AddRateOrAmount() *AmountOrRate1Choice {
	s.RateOrAmount = new(AmountOrRate1Choice)
	return s.RateOrAmount
}
