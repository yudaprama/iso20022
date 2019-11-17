package model

// Elements used to calculate the collateral margin call for the variation margin.
type VariationMargin1 struct {

	// Amount of unsecured exposure a counterparty will accept before issuing a margin call in the base currency.
	ThresholdAmount *ActiveCurrencyAndAmount `xml:"ThrshldAmt"`

	// Specifies if the threshold amount is secured or unsecured.
	ThresholdType *ThresholdType1Code `xml:"ThrshldTp,omitempty"`

	// Minimum amount to pay/receive as specified in the agreement in the base currency (to avoid the need to transfer an inconveniently small amount of variation margin).
	MinimumTransferAmount *ActiveCurrencyAndAmount `xml:"MinTrfAmt"`

	// Amount specified to avoid the need to transfer uneven amounts of collateral.
	RoundingAmount *ActiveCurrencyAndAmount `xml:"RndgAmt"`

	// Defines how the rounding amount was applied in the calculation. For example, should the amount of collateral required be rounded up, down, to the closer integral multiple specified or not rounded.
	RoundingMethod *RoundingMethod1Code `xml:"RndgMtd"`
}

func (v *VariationMargin1) SetThresholdAmount(value, currency string) {
	v.ThresholdAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (v *VariationMargin1) SetThresholdType(value string) {
	v.ThresholdType = (*ThresholdType1Code)(&value)
}

func (v *VariationMargin1) SetMinimumTransferAmount(value, currency string) {
	v.MinimumTransferAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (v *VariationMargin1) SetRoundingAmount(value, currency string) {
	v.RoundingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (v *VariationMargin1) SetRoundingMethod(value string) {
	v.RoundingMethod = (*RoundingMethod1Code)(&value)
}
