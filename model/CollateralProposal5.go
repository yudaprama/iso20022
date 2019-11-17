package model

// Provides details about the proposal for the variation margin and optionaly the segregated independent amount.
type CollateralProposal5 struct {

	// Provides details about the proposal for the variation margin.
	VariationMargin *CollateralMovement7 `xml:"VartnMrgn"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement7 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposal5) AddVariationMargin() *CollateralMovement7 {
	c.VariationMargin = new(CollateralMovement7)
	return c.VariationMargin
}

func (c *CollateralProposal5) AddSegregatedIndependentAmount() *CollateralMovement7 {
	c.SegregatedIndependentAmount = new(CollateralMovement7)
	return c.SegregatedIndependentAmount
}
