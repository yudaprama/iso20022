package model

// Indicates the type of proposal and if the proposal is  for the variation margin and the segregated independent amount, or the segregated independent amount only.
type Proposal5 struct {

	// Indicates whether this is an initial or counter proposal.
	CollateralProposalType *ProposalType1Code `xml:"CollPrpslTp"`

	// Provides details about the proposal for the variation margin and the segregated independent amount, or the segregated independent amount only.
	CollateralProposal *CollateralProposal5Choice `xml:"CollPrpsl"`
}

func (p *Proposal5) SetCollateralProposalType(value string) {
	p.CollateralProposalType = (*ProposalType1Code)(&value)
}

func (p *Proposal5) AddCollateralProposal() *CollateralProposal5Choice {
	p.CollateralProposal = new(CollateralProposal5Choice)
	return p.CollateralProposal
}
