package model

// Information related to a linked transaction.
type Reference17 struct {

	// Identification of the collateral substitution request.
	CollateralSubstitutionRequestIdentification *Max35Text `xml:"CollSbstitnReqId"`

	// Identification of the collateral substitution response.
	CollateralSubstitutionResponseIdentification *Max35Text `xml:"CollSbstitnRspnId,omitempty"`
}

func (r *Reference17) SetCollateralSubstitutionRequestIdentification(value string) {
	r.CollateralSubstitutionRequestIdentification = (*Max35Text)(&value)
}

func (r *Reference17) SetCollateralSubstitutionResponseIdentification(value string) {
	r.CollateralSubstitutionResponseIdentification = (*Max35Text)(&value)
}
