package model

// Provides the collateral proposal response for the variation margin and the segregated independent amount, or the segregated independent amount only.
type CollateralProposalResponse2Choice struct {

	// Provides the collateral proposal response for the variation margin and optionaly the segregated independent amount.
	CollateralProposal *CollateralProposalResponse2 `xml:"CollPrpsl"`

	// Provides the collateral proposal response for the segregated independent amount only.
	SegregatedIndependentAmount *CollateralProposalResponseType2 `xml:"SgrtdIndpdntAmt"`
}

func (c *CollateralProposalResponse2Choice) AddCollateralProposal() *CollateralProposalResponse2 {
	c.CollateralProposal = new(CollateralProposalResponse2)
	return c.CollateralProposal
}

func (c *CollateralProposalResponse2Choice) AddSegregatedIndependentAmount() *CollateralProposalResponseType2 {
	c.SegregatedIndependentAmount = new(CollateralProposalResponseType2)
	return c.SegregatedIndependentAmount
}
