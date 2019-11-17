package model

// Provides details about the proposal for the variation margin and optionaly the segregated independent amount, or the segregated independent amount only.
type CollateralProposal4Choice struct {

	// Provides details about the proposal for the variation margin and optionaly the segregated independent amount.
	CollateralProposalDetails *CollateralProposal5 `xml:"CollPrpslDtls"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement7 `xml:"SgrtdIndpdntAmt"`
}

func (c *CollateralProposal4Choice) AddCollateralProposalDetails() *CollateralProposal5 {
	c.CollateralProposalDetails = new(CollateralProposal5)
	return c.CollateralProposalDetails
}

func (c *CollateralProposal4Choice) AddSegregatedIndependentAmount() *CollateralMovement7 {
	c.SegregatedIndependentAmount = new(CollateralMovement7)
	return c.SegregatedIndependentAmount
}
