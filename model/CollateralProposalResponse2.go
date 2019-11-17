package model

// Provides the collateral proposal response for the variation margin and optionaly the segregated independent amount.
type CollateralProposalResponse2 struct {

	// Provides the collateral proposal response for the variation margin.
	VariationMargin *CollateralProposalResponseType2 `xml:"VartnMrgn"`

	// Provides the collateral proposal response for the segregated independent amount.
	SegregatedIndependentAmount *CollateralProposalResponseType2 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposalResponse2) AddVariationMargin() *CollateralProposalResponseType2 {
	c.VariationMargin = new(CollateralProposalResponseType2)
	return c.VariationMargin
}

func (c *CollateralProposalResponse2) AddSegregatedIndependentAmount() *CollateralProposalResponseType2 {
	c.SegregatedIndependentAmount = new(CollateralProposalResponseType2)
	return c.SegregatedIndependentAmount
}
