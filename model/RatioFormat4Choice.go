package model

// Choice of format to expressed a ratio.
type RatioFormat4Choice struct {

	// Ratio expressed as a quotient of quantities.
	QuantityToQuantity *QuantityToQuantityRatio1 `xml:"QtyToQty"`

	// Ratio expressed as a quotient of amounts.
	AmountToAmount *AmountToAmountRatio2 `xml:"AmtToAmt"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio2 `xml:"AmtToQty"`

	// Ratio expressed as a quantity to amount ratio.
	QuantityToAmount *AmountAndQuantityRatio2 `xml:"QtyToAmt"`
}

func (r *RatioFormat4Choice) AddQuantityToQuantity() *QuantityToQuantityRatio1 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio1)
	return r.QuantityToQuantity
}

func (r *RatioFormat4Choice) AddAmountToAmount() *AmountToAmountRatio2 {
	r.AmountToAmount = new(AmountToAmountRatio2)
	return r.AmountToAmount
}

func (r *RatioFormat4Choice) AddAmountToQuantity() *AmountAndQuantityRatio2 {
	r.AmountToQuantity = new(AmountAndQuantityRatio2)
	return r.AmountToQuantity
}

func (r *RatioFormat4Choice) AddQuantityToAmount() *AmountAndQuantityRatio2 {
	r.QuantityToAmount = new(AmountAndQuantityRatio2)
	return r.QuantityToAmount
}
