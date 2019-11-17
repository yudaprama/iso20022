package model

// Provides details on the response for a collateral proposal.
type CollateralProposalResponseType3 struct {

	// Unique identifier for a collateral proposal.
	CollateralProposalIdentification *Max35Text `xml:"CollPrpslId"`

	// Indicates whether the collateral proposal is an initial or a counter proposal.
	Type *CollateralProposalResponse1Code `xml:"Tp"`

	// Provides response details for each of the proposed collateral pieces.
	Response *CollateralResponse2 `xml:"Rspn"`
}

func (c *CollateralProposalResponseType3) SetCollateralProposalIdentification(value string) {
	c.CollateralProposalIdentification = (*Max35Text)(&value)
}

func (c *CollateralProposalResponseType3) SetType(value string) {
	c.Type = (*CollateralProposalResponse1Code)(&value)
}

func (c *CollateralProposalResponseType3) AddResponse() *CollateralResponse2 {
	c.Response = new(CollateralResponse2)
	return c.Response
}
