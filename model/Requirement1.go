package model

// Provides details about the margin requirements for the variation margin and optionally the segregated independent amount.
type Requirement1 struct {

	// Provides details about the margin requirements for the variation margin.
	VariationMarginRequirement *MarginRequirement1 `xml:"VartnMrgnRqrmnt"`

	// Provides details about the margin requirements for the segregated independent amount.
	SegregatedIndependentAmountRequirement *MarginRequirement1 `xml:"SgrtdIndpdntAmtRqrmnt,omitempty"`
}

func (r *Requirement1) AddVariationMarginRequirement() *MarginRequirement1 {
	r.VariationMarginRequirement = new(MarginRequirement1)
	return r.VariationMarginRequirement
}

func (r *Requirement1) AddSegregatedIndependentAmountRequirement() *MarginRequirement1 {
	r.SegregatedIndependentAmountRequirement = new(MarginRequirement1)
	return r.SegregatedIndependentAmountRequirement
}
