package model

// Choice between a rate or an unspecified rate.
type SolicitationFeeRateFormat10Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio5 `xml:"AmtToQty"`

	// Cash amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (s *SolicitationFeeRateFormat10Choice) SetRate(value string) {
	s.Rate = (*PercentageRate)(&value)
}

func (s *SolicitationFeeRateFormat10Choice) AddAmountToQuantity() *AmountAndQuantityRatio5 {
	s.AmountToQuantity = new(AmountAndQuantityRatio5)
	return s.AmountToQuantity
}

func (s *SolicitationFeeRateFormat10Choice) SetAmount(value, currency string) {
	s.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (s *SolicitationFeeRateFormat10Choice) SetNotSpecifiedRate(value string) {
	s.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
