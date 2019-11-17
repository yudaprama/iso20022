package model

// Provides details about the proposal for the variation margin and optionaly the segregated independent amount.
type CollateralProposal4 struct {

	// Provides details about the proposal for the variation margin.
	VariationMargin *CollateralMovement5 `xml:"VartnMrgn"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement5 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposal4) AddVariationMargin() *CollateralMovement5 {
	c.VariationMargin = new(CollateralMovement5)
	return c.VariationMargin
}

func (c *CollateralProposal4) AddSegregatedIndependentAmount() *CollateralMovement5 {
	c.SegregatedIndependentAmount = new(CollateralMovement5)
	return c.SegregatedIndependentAmount
}
