package model

// Choice of format to expressed a ratio.
type RatioFormat19Choice struct {

	// Ratio expressed as a quotient of quantities.
	QuantityToQuantity *QuantityToQuantityRatio1 `xml:"QtyToQty"`

	// Ratio expressed as a quotient of amounts.
	AmountToAmount *AmountToAmountRatio2 `xml:"AmtToAmt"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio4 `xml:"AmtToQty"`

	// Ratio expressed as a quantity to amount ratio.
	QuantityToAmount *AmountAndQuantityRatio4 `xml:"QtyToAmt"`
}

func (r *RatioFormat19Choice) AddQuantityToQuantity() *QuantityToQuantityRatio1 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio1)
	return r.QuantityToQuantity
}

func (r *RatioFormat19Choice) AddAmountToAmount() *AmountToAmountRatio2 {
	r.AmountToAmount = new(AmountToAmountRatio2)
	return r.AmountToAmount
}

func (r *RatioFormat19Choice) AddAmountToQuantity() *AmountAndQuantityRatio4 {
	r.AmountToQuantity = new(AmountAndQuantityRatio4)
	return r.AmountToQuantity
}

func (r *RatioFormat19Choice) AddQuantityToAmount() *AmountAndQuantityRatio4 {
	r.QuantityToAmount = new(AmountAndQuantityRatio4)
	return r.QuantityToAmount
}
