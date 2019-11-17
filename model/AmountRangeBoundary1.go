package model

// Limit for an amount range.
type AmountRangeBoundary1 struct {

	// Amount value of the range limit.
	BoundaryAmount *ImpliedCurrencyAndAmount `xml:"BdryAmt"`

	// Indicates whether the boundary amount is included in the range of amount values.
	Included *YesNoIndicator `xml:"Incl"`
}

func (a *AmountRangeBoundary1) SetBoundaryAmount(value, currency string) {
	a.BoundaryAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *AmountRangeBoundary1) SetIncluded(value string) {
	a.Included = (*YesNoIndicator)(&value)
}
