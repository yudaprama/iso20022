package model

// Choice between a rate or an unspecified rate.
type SolicitationFeeRateFormat9Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio5 `xml:"AmtToQty"`

	// Cash amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (s *SolicitationFeeRateFormat9Choice) SetRate(value string) {
	s.Rate = (*PercentageRate)(&value)
}

func (s *SolicitationFeeRateFormat9Choice) AddAmountToQuantity() *AmountAndQuantityRatio5 {
	s.AmountToQuantity = new(AmountAndQuantityRatio5)
	return s.AmountToQuantity
}

func (s *SolicitationFeeRateFormat9Choice) SetAmount(value, currency string) {
	s.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
