package model

// Ratio expressed as a quotient of quantities.
type QuantityToQuantityRatio1 struct {

	// Numerator of the quotient of quantities.
	Quantity1 *DecimalNumber `xml:"Qty1"`

	// Denominator of the quotient of quantities.
	Quantity2 *DecimalNumber `xml:"Qty2"`
}

func (q *QuantityToQuantityRatio1) SetQuantity1(value string) {
	q.Quantity1 = (*DecimalNumber)(&value)
}

func (q *QuantityToQuantityRatio1) SetQuantity2(value string) {
	q.Quantity2 = (*DecimalNumber)(&value)
}
