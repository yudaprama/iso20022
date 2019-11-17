package model

// Choice between an amount and percentage.
type AmountOrPercentage1Choice struct {

	// Details related to a defined monetary amount.
	DefinedAmount *UndertakingAmount4 `xml:"DfndAmt"`

	// Details related to an amount percentage.
	PercentageAmount *Percentage1 `xml:"PctgAmt"`
}

func (a *AmountOrPercentage1Choice) AddDefinedAmount() *UndertakingAmount4 {
	a.DefinedAmount = new(UndertakingAmount4)
	return a.DefinedAmount
}

func (a *AmountOrPercentage1Choice) AddPercentageAmount() *Percentage1 {
	a.PercentageAmount = new(Percentage1)
	return a.PercentageAmount
}
