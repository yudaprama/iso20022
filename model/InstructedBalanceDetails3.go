package model

// Provides information about total instructed balance.
type InstructedBalanceDetails3 struct {

	// Provides information about the total instructed balance.
	TotalInstructedBalance *BalanceFormat1Choice `xml:"TtlInstdBal"`

	// Provide instructed balance breakdown information per option.
	OptionDetails []*InstructedCorporateActionOption4 `xml:"OptnDtls,omitempty"`
}

func (i *InstructedBalanceDetails3) AddTotalInstructedBalance() *BalanceFormat1Choice {
	i.TotalInstructedBalance = new(BalanceFormat1Choice)
	return i.TotalInstructedBalance
}

func (i *InstructedBalanceDetails3) AddOptionDetails() *InstructedCorporateActionOption4 {
	newValue := new(InstructedCorporateActionOption4)
	i.OptionDetails = append(i.OptionDetails, newValue)
	return newValue
}
