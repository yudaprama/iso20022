package model

// Provides details about the collateral against variation margin and optionally the segregated independent amount.
type Collateral1 struct {

	// Provides details about the collateral held, in transit or that still needs to be agreed by both parties, against the variation margin.
	VariationMargin *MarginCollateral1 `xml:"VartnMrgn"`

	// Provides details about the collateral held, in transit or that still needs to be agreed by both parties, against the segregated independent amount.
	SegregatedIndependentAmount *MarginCollateral1 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *Collateral1) AddVariationMargin() *MarginCollateral1 {
	c.VariationMargin = new(MarginCollateral1)
	return c.VariationMargin
}

func (c *Collateral1) AddSegregatedIndependentAmount() *MarginCollateral1 {
	c.SegregatedIndependentAmount = new(MarginCollateral1)
	return c.SegregatedIndependentAmount
}
