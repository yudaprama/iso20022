package model

// Choice between a rate or an unspecified rate.
type SolicitationFeeRateFormat1Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType6Code `xml:"NotSpcfdRate"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio2 `xml:"AmtToQty"`
}

func (s *SolicitationFeeRateFormat1Choice) SetRate(value string) {
	s.Rate = (*PercentageRate)(&value)
}

func (s *SolicitationFeeRateFormat1Choice) SetNotSpecifiedRate(value string) {
	s.NotSpecifiedRate = (*RateValueType6Code)(&value)
}

func (s *SolicitationFeeRateFormat1Choice) AddAmountToQuantity() *AmountAndQuantityRatio2 {
	s.AmountToQuantity = new(AmountAndQuantityRatio2)
	return s.AmountToQuantity
}
