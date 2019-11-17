package model

// Provides the margin requirements for the variation margin and the segregated independent amount, or the segregated independent amount only.
type MarginRequirement1Choice struct {

	// Provides details about the margin requirements for the variation margin and optionally the segregated independent amount.
	MarginRequirement *Requirement1 `xml:"MrgnRqrmnt"`

	// Provides details about the margin requirements for the segregated independent amount only.
	SegregatedIndependentAmountRequirement *MarginRequirement1 `xml:"SgrtdIndpdntAmtRqrmnt"`
}

func (m *MarginRequirement1Choice) AddMarginRequirement() *Requirement1 {
	m.MarginRequirement = new(Requirement1)
	return m.MarginRequirement
}

func (m *MarginRequirement1Choice) AddSegregatedIndependentAmountRequirement() *MarginRequirement1 {
	m.SegregatedIndependentAmountRequirement = new(MarginRequirement1)
	return m.SegregatedIndependentAmountRequirement
}
