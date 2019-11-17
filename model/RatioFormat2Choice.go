package model

// Choice of format to expressed a ratio.
type RatioFormat2Choice struct {

	// The ratio is expressed as a quantity per another quantity.
	QuantityToQuantity *QuantityToQuantityRatio1 `xml:"QtyToQty"`

	// The ratio is expressed as an amount per another amount.
	AmountToAmount *AmountToAmountRatio1 `xml:"AmtToAmt"`

	// The ratio is expressed as an amount per quantity.
	AmountToQuantity *AmountAndQuantityRatio1 `xml:"AmtToQty"`

	// Ratio is expressed as a quantity per an amount.
	QuantityToAmount *AmountAndQuantityRatio1 `xml:"QtyToAmt"`

	// The value of the ratio is not specified.
	NotSpecifiedRate *RateType12FormatChoice `xml:"NotSpcfdRate"`
}

func (r *RatioFormat2Choice) AddQuantityToQuantity() *QuantityToQuantityRatio1 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio1)
	return r.QuantityToQuantity
}

func (r *RatioFormat2Choice) AddAmountToAmount() *AmountToAmountRatio1 {
	r.AmountToAmount = new(AmountToAmountRatio1)
	return r.AmountToAmount
}

func (r *RatioFormat2Choice) AddAmountToQuantity() *AmountAndQuantityRatio1 {
	r.AmountToQuantity = new(AmountAndQuantityRatio1)
	return r.AmountToQuantity
}

func (r *RatioFormat2Choice) AddQuantityToAmount() *AmountAndQuantityRatio1 {
	r.QuantityToAmount = new(AmountAndQuantityRatio1)
	return r.QuantityToAmount
}

func (r *RatioFormat2Choice) AddNotSpecifiedRate() *RateType12FormatChoice {
	r.NotSpecifiedRate = new(RateType12FormatChoice)
	return r.NotSpecifiedRate
}
