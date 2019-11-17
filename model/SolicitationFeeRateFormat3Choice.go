package model

// Choice between a rate or an unspecified rate.
type SolicitationFeeRateFormat3Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio2 `xml:"AmtToQty"`
}

func (s *SolicitationFeeRateFormat3Choice) SetRate(value string) {
	s.Rate = (*PercentageRate)(&value)
}

func (s *SolicitationFeeRateFormat3Choice) SetNotSpecifiedRate(value string) {
	s.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (s *SolicitationFeeRateFormat3Choice) AddAmountToQuantity() *AmountAndQuantityRatio2 {
	s.AmountToQuantity = new(AmountAndQuantityRatio2)
	return s.AmountToQuantity
}
