package model

// Provides information about total instructed balance.
type InstructedBalanceDetails6 struct {

	// Provides information about the total instructed balance.
	TotalInstructedBalance *BalanceFormat7Choice `xml:"TtlInstdBal"`

	// Provide instructed balance breakdown information per option.
	OptionDetails []*InstructedCorporateActionOption7 `xml:"OptnDtls,omitempty"`
}

func (i *InstructedBalanceDetails6) AddTotalInstructedBalance() *BalanceFormat7Choice {
	i.TotalInstructedBalance = new(BalanceFormat7Choice)
	return i.TotalInstructedBalance
}

func (i *InstructedBalanceDetails6) AddOptionDetails() *InstructedCorporateActionOption7 {
	newValue := new(InstructedCorporateActionOption7)
	i.OptionDetails = append(i.OptionDetails, newValue)
	return newValue
}
