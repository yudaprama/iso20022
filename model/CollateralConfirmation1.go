package model

// Provides the status details about the collateral substitution.
type CollateralConfirmation1 struct {

	// Reference to the collateral substitution request identification.
	CollateralSubstitutionRequestIdentification *Max35Text `xml:"CollSbstitnReqId"`

	// Reference to the collateral substitution response identification.
	CollateralSubstitutionResponseIdentification *Max35Text `xml:"CollSbstitnRspnId,omitempty"`

	// Provides details about the status of the collateral substitution, either released or returned.
	ConfirmationType *CollateralSubstitutionConfirmation1Code `xml:"ConfTp"`

	// Allows to provides additional comments on the collateral substitution status.
	Comment *Max140Text `xml:"Cmnt,omitempty"`
}

func (c *CollateralConfirmation1) SetCollateralSubstitutionRequestIdentification(value string) {
	c.CollateralSubstitutionRequestIdentification = (*Max35Text)(&value)
}

func (c *CollateralConfirmation1) SetCollateralSubstitutionResponseIdentification(value string) {
	c.CollateralSubstitutionResponseIdentification = (*Max35Text)(&value)
}

func (c *CollateralConfirmation1) SetConfirmationType(value string) {
	c.ConfirmationType = (*CollateralSubstitutionConfirmation1Code)(&value)
}

func (c *CollateralConfirmation1) SetComment(value string) {
	c.Comment = (*Max140Text)(&value)
}
