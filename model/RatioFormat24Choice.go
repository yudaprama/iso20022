package model

// Choice of format to express a ratio.
type RatioFormat24Choice struct {

	// Ratio expressed as a quotient of quantities.
	QuantityToQuantity *QuantityToQuantityRatio2 `xml:"QtyToQty"`

	// Value of the ratio not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Ratio expressed as a quotient of amounts.
	AmountToAmount *AmountToAmountRatio3 `xml:"AmtToAmt"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio5 `xml:"AmtToQty"`

	// Ratio expressed as a quantity to amount ratio.
	QuantityToAmount *AmountAndQuantityRatio5 `xml:"QtyToAmt"`
}

func (r *RatioFormat24Choice) AddQuantityToQuantity() *QuantityToQuantityRatio2 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio2)
	return r.QuantityToQuantity
}

func (r *RatioFormat24Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RatioFormat24Choice) AddAmountToAmount() *AmountToAmountRatio3 {
	r.AmountToAmount = new(AmountToAmountRatio3)
	return r.AmountToAmount
}

func (r *RatioFormat24Choice) AddAmountToQuantity() *AmountAndQuantityRatio5 {
	r.AmountToQuantity = new(AmountAndQuantityRatio5)
	return r.AmountToQuantity
}

func (r *RatioFormat24Choice) AddQuantityToAmount() *AmountAndQuantityRatio5 {
	r.QuantityToAmount = new(AmountAndQuantityRatio5)
	return r.QuantityToAmount
}
