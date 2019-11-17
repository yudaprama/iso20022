package model

// Choice of format to express a ratio.
type RatioFormat23Choice struct {

	// Ratio expressed as a quotient of quantities.
	QuantityToQuantity *QuantityToQuantityRatio2 `xml:"QtyToQty"`

	// Value of the ratio not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Ratio expressed as a quotient of amounts.
	AmountToAmount *AmountToAmountRatio3 `xml:"AmtToAmt"`
}

func (r *RatioFormat23Choice) AddQuantityToQuantity() *QuantityToQuantityRatio2 {
	r.QuantityToQuantity = new(QuantityToQuantityRatio2)
	return r.QuantityToQuantity
}

func (r *RatioFormat23Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RatioFormat23Choice) AddAmountToAmount() *AmountToAmountRatio3 {
	r.AmountToAmount = new(AmountToAmountRatio3)
	return r.AmountToAmount
}
