package model

// Choice between a rate or an unspecified rate.
type SolicitationFeeRateFormat5Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio2 `xml:"AmtToQty"`

	// Cash amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (s *SolicitationFeeRateFormat5Choice) SetRate(value string) {
	s.Rate = (*PercentageRate)(&value)
}

func (s *SolicitationFeeRateFormat5Choice) AddAmountToQuantity() *AmountAndQuantityRatio2 {
	s.AmountToQuantity = new(AmountAndQuantityRatio2)
	return s.AmountToQuantity
}

func (s *SolicitationFeeRateFormat5Choice) SetAmount(value, currency string) {
	s.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (s *SolicitationFeeRateFormat5Choice) SetNotSpecifiedRate(value string) {
	s.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
