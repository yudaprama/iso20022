package model

// Elements used to calculate the collateral margin call for the segregated independent amount.
type SegregatedIndependentAmountMargin1 struct {

	// Minimum amount to pay/receive as specified in the agreement in the base currency (to avoid the need to transfer an inconveniently small amount of segregated independent amount).
	MinimumTransferAmount *ActiveCurrencyAndAmount `xml:"MinTrfAmt"`

	// Amount specified to avoid the need to transfer uneven amounts of independent amount collateral.
	RoundingAmount *ActiveCurrencyAndAmount `xml:"RndgAmt,omitempty"`

	// Defines how the rounding amount was applied in the calculation. For example, should the amount of collateral required be rounded up, down, to the closer integral multiple specified or not rounded.
	RoundingMethod *RoundingMethod1Code `xml:"RndgMtd,omitempty"`
}

func (s *SegregatedIndependentAmountMargin1) SetMinimumTransferAmount(value, currency string) {
	s.MinimumTransferAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SegregatedIndependentAmountMargin1) SetRoundingAmount(value, currency string) {
	s.RoundingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SegregatedIndependentAmountMargin1) SetRoundingMethod(value string) {
	s.RoundingMethod = (*RoundingMethod1Code)(&value)
}
