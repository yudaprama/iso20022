package model

// Choice of format to expressed a ratio.
type RatioFormat1Choice struct {

	// The ratio is expressed as a quantity per another quantity.
	QuantityToQuantity *QuantityToQuantityRatio1 `xml:"QtyToQty"`

	// The ratio is expressed as an amount per another amount.
	AmountToAmount *AmountToAmountRatio1 `xml:"AmtToAmt"`

	// The value of the ratio is not specified.
	NotSpecifiedRate *RateType12FormatChoice `xml:"NotSpcfdRate"`
}

func (r *RatioFormat1Choice) AddQuantityToQuantity() *QuantityToQuantityRatio1 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio1)
	return r.QuantityToQuantity
}

func (r *RatioFormat1Choice) AddAmountToAmount() *AmountToAmountRatio1 {
	r.AmountToAmount = new(AmountToAmountRatio1)
	return r.AmountToAmount
}

func (r *RatioFormat1Choice) AddNotSpecifiedRate() *RateType12FormatChoice {
	r.NotSpecifiedRate = new(RateType12FormatChoice)
	return r.NotSpecifiedRate
}
