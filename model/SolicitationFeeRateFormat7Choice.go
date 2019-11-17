package model

// Choice between a rate or an unspecified rate.
type SolicitationFeeRateFormat7Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio4 `xml:"AmtToQty"`

	// Cash amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (s *SolicitationFeeRateFormat7Choice) SetRate(value string) {
	s.Rate = (*PercentageRate)(&value)
}

func (s *SolicitationFeeRateFormat7Choice) AddAmountToQuantity() *AmountAndQuantityRatio4 {
	s.AmountToQuantity = new(AmountAndQuantityRatio4)
	return s.AmountToQuantity
}

func (s *SolicitationFeeRateFormat7Choice) SetAmount(value, currency string) {
	s.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (s *SolicitationFeeRateFormat7Choice) SetNotSpecifiedRate(value string) {
	s.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
