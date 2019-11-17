package model

// Choice between a rate or an unspecified rate.
type SolicitationFeeRateFormat8Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio4 `xml:"AmtToQty"`

	// Cash amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (s *SolicitationFeeRateFormat8Choice) SetRate(value string) {
	s.Rate = (*PercentageRate)(&value)
}

func (s *SolicitationFeeRateFormat8Choice) AddAmountToQuantity() *AmountAndQuantityRatio4 {
	s.AmountToQuantity = new(AmountAndQuantityRatio4)
	return s.AmountToQuantity
}

func (s *SolicitationFeeRateFormat8Choice) SetAmount(value, currency string) {
	s.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
