package model

// Provides details about the proposal for the variation margin and optionally the segregated independent amount, or the segregated independent amount only.
type CollateralProposal5Choice struct {

	// Provides details about the proposal for the variation margin and optionally the segregated independent amount.
	CollateralProposalDetails *CollateralProposal6 `xml:"CollPrpslDtls"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement10 `xml:"SgrtdIndpdntAmt"`
}

func (c *CollateralProposal5Choice) AddCollateralProposalDetails() *CollateralProposal6 {
	c.CollateralProposalDetails = new(CollateralProposal6)
	return c.CollateralProposalDetails
}

func (c *CollateralProposal5Choice) AddSegregatedIndependentAmount() *CollateralMovement10 {
	c.SegregatedIndependentAmount = new(CollateralMovement10)
	return c.SegregatedIndependentAmount
}
