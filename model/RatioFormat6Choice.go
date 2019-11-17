package model

// Choice of format to expressed a ratio.
type RatioFormat6Choice struct {

	// Ratio expressed as a quotient of quantities.
	QuantityToQuantity *QuantityToQuantityRatio1 `xml:"QtyToQty"`

	// Value of the ratio not specified.
	NotSpecifiedRate *RateValueType6Code `xml:"NotSpcfdRate"`

	// Ratio expressed as a quotient of amounts.
	AmountToAmount *AmountToAmountRatio2 `xml:"AmtToAmt"`

	// Ratio expressed as an amount to quantity ratio.
	AmountToQuantity *AmountAndQuantityRatio2 `xml:"AmtToQty"`

	// Ratio expressed as a quantity to amount ratio.
	QuantityToAmount *AmountAndQuantityRatio2 `xml:"QtyToAmt"`
}

func (r *RatioFormat6Choice) AddQuantityToQuantity() *QuantityToQuantityRatio1 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio1)
	return r.QuantityToQuantity
}

func (r *RatioFormat6Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType6Code)(&value)
}

func (r *RatioFormat6Choice) AddAmountToAmount() *AmountToAmountRatio2 {
	r.AmountToAmount = new(AmountToAmountRatio2)
	return r.AmountToAmount
}

func (r *RatioFormat6Choice) AddAmountToQuantity() *AmountAndQuantityRatio2 {
	r.AmountToQuantity = new(AmountAndQuantityRatio2)
	return r.AmountToQuantity
}

func (r *RatioFormat6Choice) AddQuantityToAmount() *AmountAndQuantityRatio2 {
	r.QuantityToAmount = new(AmountAndQuantityRatio2)
	return r.QuantityToAmount
}
