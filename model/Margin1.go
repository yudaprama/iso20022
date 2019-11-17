package model

// Defines the elements used to calculate the collateral margin call for the variation margin and optionally the segregated independent amount.
type Margin1 struct {

	// Elements used to calculate the collateral margin call for the variation margin.
	VariationMargin *VariationMargin1 `xml:"VartnMrgn"`

	// Elements used to calculate the collateral margin call for the segregated independent amount.
	SegregatedIndependentAmountMargin *SegregatedIndependentAmountMargin1 `xml:"SgrtdIndpdntAmtMrgn,omitempty"`
}

func (m *Margin1) AddVariationMargin() *VariationMargin1 {
	m.VariationMargin = new(VariationMargin1)
	return m.VariationMargin
}

func (m *Margin1) AddSegregatedIndependentAmountMargin() *SegregatedIndependentAmountMargin1 {
	m.SegregatedIndependentAmountMargin = new(SegregatedIndependentAmountMargin1)
	return m.SegregatedIndependentAmountMargin
}
