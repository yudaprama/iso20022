package model

// Choice to provide the summation of the call amounts for the variation margin and the segregated independent amount, or the segregated independent amount only.
type MarginCallResult2Choice struct {

	// Provides the summation of the call amounts for the variation margin and optionaly the segregated independent amount.
	MarginCallResultDetails *MarginCallResult2 `xml:"MrgnCallRsltDtls"`

	// Provides the summation of the call amounts.
	MarginCallAmount *Result1 `xml:"MrgnCallAmt"`

	// Provides the summation of the call amounts for the segregated independent amount only.
	SegregatedIndependentAmount *Result1 `xml:"SgrtdIndpdntAmt"`
}

func (m *MarginCallResult2Choice) AddMarginCallResultDetails() *MarginCallResult2 {
	m.MarginCallResultDetails = new(MarginCallResult2)
	return m.MarginCallResultDetails
}

func (m *MarginCallResult2Choice) AddMarginCallAmount() *Result1 {
	m.MarginCallAmount = new(Result1)
	return m.MarginCallAmount
}

func (m *MarginCallResult2Choice) AddSegregatedIndependentAmount() *Result1 {
	m.SegregatedIndependentAmount = new(Result1)
	return m.SegregatedIndependentAmount
}
