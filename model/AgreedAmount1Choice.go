package model

// Provides details about the agreed amount for the variation margin and the segregated independent amount, or the segregated independent amount only.
type AgreedAmount1Choice struct {

	// Provides details about the agreed amount for the variation margin and optionaly the segregated independent amount.
	AgreedAmountDetails *AgreedAmount1 `xml:"AgrdAmtDtls"`

	// Provides details about the agreed amount for the segregated independent amount.
	SegregatedIndependentAmount *Amount1 `xml:"SgrtdIndpdntAmt"`
}

func (a *AgreedAmount1Choice) AddAgreedAmountDetails() *AgreedAmount1 {
	a.AgreedAmountDetails = new(AgreedAmount1)
	return a.AgreedAmountDetails
}

func (a *AgreedAmount1Choice) AddSegregatedIndependentAmount() *Amount1 {
	a.SegregatedIndependentAmount = new(Amount1)
	return a.SegregatedIndependentAmount
}
