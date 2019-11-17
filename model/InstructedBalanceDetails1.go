package model

// Provides information about total instructed balance.
type InstructedBalanceDetails1 struct {

	// Provides information about the total instructed balance.
	TotalInstructedBalance *BalanceFormat1Choice `xml:"TtlInstdBal"`

	// Provide instructed balance breakdown information per option.
	OptionDetails []*InstructedCorporateActionOption1 `xml:"OptnDtls,omitempty"`
}

func (i *InstructedBalanceDetails1) AddTotalInstructedBalance() *BalanceFormat1Choice {
	i.TotalInstructedBalance = new(BalanceFormat1Choice)
	return i.TotalInstructedBalance
}

func (i *InstructedBalanceDetails1) AddOptionDetails() *InstructedCorporateActionOption1 {
	newValue := new(InstructedCorporateActionOption1)
	i.OptionDetails = append(i.OptionDetails, newValue)
	return newValue
}
