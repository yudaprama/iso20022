package model

// Provides the collateral proposal response for the variation margin and the segregated independent amount, or the segregated independent amount only.
type CollateralProposalResponse3Choice struct {

	// Provides the collateral proposal response for the variation margin and optionally the segregated independent amount.
	CollateralProposal *CollateralProposalResponse3 `xml:"CollPrpsl"`

	// Provides the collateral proposal response for the segregated independent amount only.
	SegregatedIndependentAmount *CollateralProposalResponseType3 `xml:"SgrtdIndpdntAmt"`
}

func (c *CollateralProposalResponse3Choice) AddCollateralProposal() *CollateralProposalResponse3 {
	c.CollateralProposal = new(CollateralProposalResponse3)
	return c.CollateralProposal
}

func (c *CollateralProposalResponse3Choice) AddSegregatedIndependentAmount() *CollateralProposalResponseType3 {
	c.SegregatedIndependentAmount = new(CollateralProposalResponseType3)
	return c.SegregatedIndependentAmount
}
