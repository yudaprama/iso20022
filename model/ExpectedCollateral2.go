package model

// Provides the expected collateral type and direction for the variation margin and optionaly the segregated independent amount.
type ExpectedCollateral2 struct {

	// Provides the expected collateral type and direction for the variation margin.
	VariationMargin *ExpectedCollateralMovement2 `xml:"VartnMrgn"`

	// Provides the expected collateral type and direction for the segregated independent amount.
	SegregatedIndependentAmount *ExpectedCollateralMovement2 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (e *ExpectedCollateral2) AddVariationMargin() *ExpectedCollateralMovement2 {
	e.VariationMargin = new(ExpectedCollateralMovement2)
	return e.VariationMargin
}

func (e *ExpectedCollateral2) AddSegregatedIndependentAmount() *ExpectedCollateralMovement2 {
	e.SegregatedIndependentAmount = new(ExpectedCollateralMovement2)
	return e.SegregatedIndependentAmount
}
