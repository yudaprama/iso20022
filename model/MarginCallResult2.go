package model

// Provides the summation of the call amounts for the variation margin and optionaly the segregated independent amount.
type MarginCallResult2 struct {

	// Provides the summation of the call amounts for the variation margin amount only.
	VariationMarginResult *Result1 `xml:"VartnMrgnRslt"`

	// Provides the summation of the call amounts for the segregated independent amount.
	SegregatedIndependentAmount *Result1 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (m *MarginCallResult2) AddVariationMarginResult() *Result1 {
	m.VariationMarginResult = new(Result1)
	return m.VariationMarginResult
}

func (m *MarginCallResult2) AddSegregatedIndependentAmount() *Result1 {
	m.SegregatedIndependentAmount = new(Result1)
	return m.SegregatedIndependentAmount
}
