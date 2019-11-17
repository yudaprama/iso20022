package model

// Provides details about the collateral substitution response.
type SubstitutionResponse1 struct {

	// Indicates if the substitution request was accepted or rejected.
	ResponseType *Status4Code `xml:"RspnTp"`

	// Provides details about the accepted collateral substitution.
	CollateralSubstitutionAcceptanceDetails *CollateralSubstitutionResponse1 `xml:"CollSbstitnAccptncDtls,omitempty"`

	// Provides details about the rejected collateral substitution.
	CollateralSubstitutionRejectionDetails *CollateralSubstitutionResponse2 `xml:"CollSbstitnRjctnDtls,omitempty"`
}

func (s *SubstitutionResponse1) SetResponseType(value string) {
	s.ResponseType = (*Status4Code)(&value)
}

func (s *SubstitutionResponse1) AddCollateralSubstitutionAcceptanceDetails() *CollateralSubstitutionResponse1 {
	s.CollateralSubstitutionAcceptanceDetails = new(CollateralSubstitutionResponse1)
	return s.CollateralSubstitutionAcceptanceDetails
}

func (s *SubstitutionResponse1) AddCollateralSubstitutionRejectionDetails() *CollateralSubstitutionResponse2 {
	s.CollateralSubstitutionRejectionDetails = new(CollateralSubstitutionResponse2)
	return s.CollateralSubstitutionRejectionDetails
}
