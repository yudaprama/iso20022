package model

// Choice of format to expressed a ratio.
type RatioFormat21Choice struct {

	// Ratio expressed as a quotient of quantities.
	QuantityToQuantity *QuantityToQuantityRatio2 `xml:"QtyToQty"`

	// Ratio expressed as a quotient of amounts.
	AmountToAmount *AmountToAmountRatio3 `xml:"AmtToAmt"`
}

func (r *RatioFormat21Choice) AddQuantityToQuantity() *QuantityToQuantityRatio2 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio2)
	return r.QuantityToQuantity
}

func (r *RatioFormat21Choice) AddAmountToAmount() *AmountToAmountRatio3 {
	r.AmountToAmount = new(AmountToAmountRatio3)
	return r.AmountToAmount
}
