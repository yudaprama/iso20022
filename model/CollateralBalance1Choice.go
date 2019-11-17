package model

// Choice to provide the collateral balance for the variation margin and the segregated independent amount, or the segregated independent amount only.
type CollateralBalance1Choice struct {

	// Collateral currently received (+)/delivered (-) in the base currency. This amount is after the haircut has been applied.
	TotalCollateral *ActiveCurrencyAndAmount `xml:"TtlColl"`

	// Provides details about the collateral held, in transit or that still needs to be agreed by both parties, for the variation margin and optionally the segregated independent amount.
	CollateralDetails *Collateral1 `xml:"CollDtls"`

	// Provides details about the collateral held, in transit or that still needs to be agreed by both parties, against the segregated independent amount only.
	SegregatedIndependentAmount *MarginCollateral1 `xml:"SgrtdIndpdntAmt"`
}

func (c *CollateralBalance1Choice) SetTotalCollateral(value, currency string) {
	c.TotalCollateral = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralBalance1Choice) AddCollateralDetails() *Collateral1 {
	c.CollateralDetails = new(Collateral1)
	return c.CollateralDetails
}

func (c *CollateralBalance1Choice) AddSegregatedIndependentAmount() *MarginCollateral1 {
	c.SegregatedIndependentAmount = new(MarginCollateral1)
	return c.SegregatedIndependentAmount
}
