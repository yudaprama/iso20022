package model

// Provides information about total instructed balance.
type InstructedBalanceDetails5 struct {

	// Provides information about the total instructed balance.
	TotalInstructedBalance *BalanceFormat5Choice `xml:"TtlInstdBal"`

	// Provide instructed balance breakdown information per option.
	OptionDetails []*InstructedCorporateActionOption6 `xml:"OptnDtls,omitempty"`
}

func (i *InstructedBalanceDetails5) AddTotalInstructedBalance() *BalanceFormat5Choice {
	i.TotalInstructedBalance = new(BalanceFormat5Choice)
	return i.TotalInstructedBalance
}

func (i *InstructedBalanceDetails5) AddOptionDetails() *InstructedCorporateActionOption6 {
	newValue := new(InstructedCorporateActionOption6)
	i.OptionDetails = append(i.OptionDetails, newValue)
	return newValue
}
