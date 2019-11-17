package model

// Provides details on the response for a collateral proposal.
type CollateralProposalResponseType2 struct {

	// Unique identifier for a collateral proposal.
	CollateralProposalIdentification *Max35Text `xml:"CollPrpslId"`

	// Indicates whether the collateral proposal is an initial or a counter proposal.
	Type *CollateralProposalResponse1Code `xml:"Tp"`

	// Provides response details for each of the proposed collateral pieces.
	Response *CollateralResponse1 `xml:"Rspn"`
}

func (c *CollateralProposalResponseType2) SetCollateralProposalIdentification(value string) {
	c.CollateralProposalIdentification = (*Max35Text)(&value)
}

func (c *CollateralProposalResponseType2) SetType(value string) {
	c.Type = (*CollateralProposalResponse1Code)(&value)
}

func (c *CollateralProposalResponseType2) AddResponse() *CollateralResponse1 {
	c.Response = new(CollateralResponse1)
	return c.Response
}
