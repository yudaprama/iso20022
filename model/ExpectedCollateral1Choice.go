package model

// Provides the expected collateral type and direction for the variation margin and the segregated independent amount, or the segregated independent amount only.
type ExpectedCollateral1Choice struct {

	// Provides the expected collateral type and direction for the variation margin and optionaly the segregated independent amount.
	ExpectedCollateralDetails *ExpectedCollateral1 `xml:"XpctdCollDtls"`

	// Provides the expected collateral type and direction for the segregated independent amount.
	SegregatedIndependentAmount *ExpectedCollateralMovement1 `xml:"SgrtdIndpdntAmt"`
}

func (e *ExpectedCollateral1Choice) AddExpectedCollateralDetails() *ExpectedCollateral1 {
	e.ExpectedCollateralDetails = new(ExpectedCollateral1)
	return e.ExpectedCollateralDetails
}

func (e *ExpectedCollateral1Choice) AddSegregatedIndependentAmount() *ExpectedCollateralMovement1 {
	e.SegregatedIndependentAmount = new(ExpectedCollateralMovement1)
	return e.SegregatedIndependentAmount
}
