package model

// Provides the expected collateral type and direction for the variation margin and the segregated independent amount, or the segregated independent amount only.
type ExpectedCollateral2Choice struct {

	// Provides the expected collateral type and direction for the variation margin and optionaly the segregated independent amount.
	ExpectedCollateralDetails *ExpectedCollateral2 `xml:"XpctdCollDtls"`

	// Provides the expected collateral type and direction for the segregated independent amount.
	SegregatedIndependentAmount *ExpectedCollateralMovement2 `xml:"SgrtdIndpdntAmt"`
}

func (e *ExpectedCollateral2Choice) AddExpectedCollateralDetails() *ExpectedCollateral2 {
	e.ExpectedCollateralDetails = new(ExpectedCollateral2)
	return e.ExpectedCollateralDetails
}

func (e *ExpectedCollateral2Choice) AddSegregatedIndependentAmount() *ExpectedCollateralMovement2 {
	e.SegregatedIndependentAmount = new(ExpectedCollateralMovement2)
	return e.SegregatedIndependentAmount
}
