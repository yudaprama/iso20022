package model

// Provides details about the accepted collateral substitution.
type CollateralSubstitutionResponse1 struct {

	// Reference to the collateral substitution request identification.
	CollateralSubstitutionRequestIdentification *Max35Text `xml:"CollSbstitnReqId"`

	// Provides the accepted collateral substitution amount.
	AcceptedAmount *ActiveCurrencyAndAmount `xml:"AccptdAmt,omitempty"`
}

func (c *CollateralSubstitutionResponse1) SetCollateralSubstitutionRequestIdentification(value string) {
	c.CollateralSubstitutionRequestIdentification = (*Max35Text)(&value)
}

func (c *CollateralSubstitutionResponse1) SetAcceptedAmount(value, currency string) {
	c.AcceptedAmount = NewActiveCurrencyAndAmount(value, currency)
}
