package model

// Choice of format to expressed a ratio.
type RatioFormat20Choice struct {

	// Ratio expressed as a quotient of quantities.
	QuantityToQuantity *QuantityToQuantityRatio1 `xml:"QtyToQty"`

	// Ratio expressed as a quotient of amounts.
	AmountToAmount *AmountToAmountRatio2 `xml:"AmtToAmt"`
}

func (r *RatioFormat20Choice) AddQuantityToQuantity() *QuantityToQuantityRatio1 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio1)
	return r.QuantityToQuantity
}

func (r *RatioFormat20Choice) AddAmountToAmount() *AmountToAmountRatio2 {
	r.AmountToAmount = new(AmountToAmountRatio2)
	return r.AmountToAmount
}
