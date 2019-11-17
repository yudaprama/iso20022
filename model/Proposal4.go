package model

// Indicates the type of proposal and if the proposal is  for the variation margin and the segregated independent amount, or the segregated independent amount only.
type Proposal4 struct {

	// Indicates whether this is an initial or counter proposal.
	CollateralProposalType *ProposalType1Code `xml:"CollPrpslTp"`

	// Provides details about the proposal for the variation margin and the segregated independent amount, or the segregated independent amount only.
	CollateralProposal *CollateralProposal4Choice `xml:"CollPrpsl"`
}

func (p *Proposal4) SetCollateralProposalType(value string) {
	p.CollateralProposalType = (*ProposalType1Code)(&value)
}

func (p *Proposal4) AddCollateralProposal() *CollateralProposal4Choice {
	p.CollateralProposal = new(CollateralProposal4Choice)
	return p.CollateralProposal
}
