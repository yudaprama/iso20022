package model

// Provides details about the agreed amount for the variation margin and optionaly the segregated independent amount.
type AgreedAmount1 struct {

	// Provides details about the agreed amount for the variation margin.
	VariationMarginAmount *Amount1 `xml:"VartnMrgnAmt"`

	// Provides details about the agreed amount for the segregated independent amount.
	SegregatedIndependentAmount *Amount1 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (a *AgreedAmount1) AddVariationMarginAmount() *Amount1 {
	a.VariationMarginAmount = new(Amount1)
	return a.VariationMarginAmount
}

func (a *AgreedAmount1) AddSegregatedIndependentAmount() *Amount1 {
	a.SegregatedIndependentAmount = new(Amount1)
	return a.SegregatedIndependentAmount
}
