package model

// Provides details about the proposal for the variation margin and optionally the segregated independent amount.
type CollateralProposal6 struct {

	// Provides details about the proposal for the variation margin.
	VariationMargin *CollateralMovement10 `xml:"VartnMrgn"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement10 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposal6) AddVariationMargin() *CollateralMovement10 {
	c.VariationMargin = new(CollateralMovement10)
	return c.VariationMargin
}

func (c *CollateralProposal6) AddSegregatedIndependentAmount() *CollateralMovement10 {
	c.SegregatedIndependentAmount = new(CollateralMovement10)
	return c.SegregatedIndependentAmount
}
