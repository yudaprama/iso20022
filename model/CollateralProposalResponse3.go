package model

// Provides the collateral proposal response for the variation margin and optionally the segregated independent amount.
type CollateralProposalResponse3 struct {

	// Provides the collateral proposal response for the variation margin.
	VariationMargin *CollateralProposalResponseType3 `xml:"VartnMrgn"`

	// Provides the collateral proposal response for the segregated independent amount.
	SegregatedIndependentAmount *CollateralProposalResponseType3 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposalResponse3) AddVariationMargin() *CollateralProposalResponseType3 {
	c.VariationMargin = new(CollateralProposalResponseType3)
	return c.VariationMargin
}

func (c *CollateralProposalResponse3) AddSegregatedIndependentAmount() *CollateralProposalResponseType3 {
	c.SegregatedIndependentAmount = new(CollateralProposalResponseType3)
	return c.SegregatedIndependentAmount
}
