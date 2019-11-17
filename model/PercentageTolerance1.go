package model

// Variance allowed on a quantity or on a price.
type PercentageTolerance1 struct {

	// Variance in percentage allowed over the agreed dimension. For example, plus 10 percent.
	PlusPercent *PercentageRate `xml:"PlusPct"`

	// Variance in percentage allowed below the agreed dimension. For example, minus 10 percent.
	MinusPercent *PercentageRate `xml:"MnsPct"`
}

func (p *PercentageTolerance1) SetPlusPercent(value string) {
	p.PlusPercent = (*PercentageRate)(&value)
}

func (p *PercentageTolerance1) SetMinusPercent(value string) {
	p.MinusPercent = (*PercentageRate)(&value)
}
