package model

// Provides the margin terms for the variation margin and the segregated independent amount, or the segregated independent amount only.
type MarginTerms1Choice struct {

	// Elements used to calculate the collateral margin call for the variation margin and optionally the segregated independent amount.
	MarginDetails *Margin1 `xml:"MrgnDtls"`

	// Elements used to calculate the collateral margin call for the segregated independent amount.
	SegregatedIndependentAmountMargin *SegregatedIndependentAmountMargin1 `xml:"SgrtdIndpdntAmtMrgn"`
}

func (m *MarginTerms1Choice) AddMarginDetails() *Margin1 {
	m.MarginDetails = new(Margin1)
	return m.MarginDetails
}

func (m *MarginTerms1Choice) AddSegregatedIndependentAmountMargin() *SegregatedIndependentAmountMargin1 {
	m.SegregatedIndependentAmountMargin = new(SegregatedIndependentAmountMargin1)
	return m.SegregatedIndependentAmountMargin
}
