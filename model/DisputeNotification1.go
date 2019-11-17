package model

// Provides the dispute notification details for the variation margin and optionaly the segregated independent amount.
type DisputeNotification1 struct {

	// Provides the dispute notification details for the variation margin.
	VariationMarginDispute *VariationMarginDispute1 `xml:"VartnMrgnDspt"`

	// Provides the dispute notification details for the segregated independent amount.
	SegregatedIndependentAmountDispute *SegregatedIndependentAmountDispute1 `xml:"SgrtdIndpdntAmtDspt,omitempty"`
}

func (d *DisputeNotification1) AddVariationMarginDispute() *VariationMarginDispute1 {
	d.VariationMarginDispute = new(VariationMarginDispute1)
	return d.VariationMarginDispute
}

func (d *DisputeNotification1) AddSegregatedIndependentAmountDispute() *SegregatedIndependentAmountDispute1 {
	d.SegregatedIndependentAmountDispute = new(SegregatedIndependentAmountDispute1)
	return d.SegregatedIndependentAmountDispute
}
