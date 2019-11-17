package model

// Related financial instrument into which the security can be converted.
type UnderlyingRatio1 struct {

	// Number of held securities for the exercise.
	UnderlyingQuantityDenominator *FinancialInstrumentQuantity1Choice `xml:"UndrlygQtyDnmtr"`

	// Number of related securities for the exercise.
	UnderlyingQuantityNumerator *FinancialInstrumentQuantity1Choice `xml:"UndrlygQtyNmrtr"`

	// Related security into which the security can be converted.
	RelatedFinancialInstrumentIdentification []*SecurityIdentification14 `xml:"RltdFinInstrmId,omitempty"`
}

func (u *UnderlyingRatio1) AddUnderlyingQuantityDenominator() *FinancialInstrumentQuantity1Choice {
	u.UnderlyingQuantityDenominator = new(FinancialInstrumentQuantity1Choice)
	return u.UnderlyingQuantityDenominator
}

func (u *UnderlyingRatio1) AddUnderlyingQuantityNumerator() *FinancialInstrumentQuantity1Choice {
	u.UnderlyingQuantityNumerator = new(FinancialInstrumentQuantity1Choice)
	return u.UnderlyingQuantityNumerator
}

func (u *UnderlyingRatio1) AddRelatedFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	u.RelatedFinancialInstrumentIdentification = append(u.RelatedFinancialInstrumentIdentification, newValue)
	return newValue
}
