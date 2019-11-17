package model

// Choice of format to express a ratio.
type RatioFormat18Choice struct {

	// Ratio expressed as a quotient of quantities.
	QuantityToQuantity *QuantityToQuantityRatio1 `xml:"QtyToQty"`

	// Value of the ratio not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Ratio expressed as a quotient of amounts.
	AmountToAmount *AmountToAmountRatio2 `xml:"AmtToAmt"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio4 `xml:"AmtToQty"`

	// Ratio expressed as a quantity to amount ratio.
	QuantityToAmount *AmountAndQuantityRatio4 `xml:"QtyToAmt"`
}

func (r *RatioFormat18Choice) AddQuantityToQuantity() *QuantityToQuantityRatio1 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio1)
	return r.QuantityToQuantity
}

func (r *RatioFormat18Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RatioFormat18Choice) AddAmountToAmount() *AmountToAmountRatio2 {
	r.AmountToAmount = new(AmountToAmountRatio2)
	return r.AmountToAmount
}

func (r *RatioFormat18Choice) AddAmountToQuantity() *AmountAndQuantityRatio4 {
	r.AmountToQuantity = new(AmountAndQuantityRatio4)
	return r.AmountToQuantity
}

func (r *RatioFormat18Choice) AddQuantityToAmount() *AmountAndQuantityRatio4 {
	r.QuantityToAmount = new(AmountAndQuantityRatio4)
	return r.QuantityToAmount
}
