package model

// Range of amount values.
type FromToAmountRange struct {

	// Lower boundary of a range of amount values.
	FromAmount *AmountRangeBoundary1 `xml:"FrAmt"`

	// Upper boundary of a range of amount values.
	ToAmount *AmountRangeBoundary1 `xml:"ToAmt"`
}

func (f *FromToAmountRange) AddFromAmount() *AmountRangeBoundary1 {
	f.FromAmount = new(AmountRangeBoundary1)
	return f.FromAmount
}

func (f *FromToAmountRange) AddToAmount() *AmountRangeBoundary1 {
	f.ToAmount = new(AmountRangeBoundary1)
	return f.ToAmount
}
