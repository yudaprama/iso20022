package model

// Provides details about the proposal for the variation margin and optionaly the segregated independent amount, or the segregated independent amount only.
type CollateralProposal3Choice struct {

	// Provides details about the proposal for the variation margin and optionaly the segregated independent amount.
	CollateralProposalDetails *CollateralProposal4 `xml:"CollPrpslDtls"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement5 `xml:"SgrtdIndpdntAmt"`
}

func (c *CollateralProposal3Choice) AddCollateralProposalDetails() *CollateralProposal4 {
	c.CollateralProposalDetails = new(CollateralProposal4)
	return c.CollateralProposalDetails
}

func (c *CollateralProposal3Choice) AddSegregatedIndependentAmount() *CollateralMovement5 {
	c.SegregatedIndependentAmount = new(CollateralMovement5)
	return c.SegregatedIndependentAmount
}
