package model

// Balance related details for a portfolio.
type BalanceDetails5 struct {

	// Balance type.
	Type *BalanceType6Choice `xml:"Tp"`

	// Unrealised gain or loss.
	Unrealised *Unrealised1Code `xml:"Urlsd,omitempty"`

	// Balance amount.
	Amount *AmountAndDirection31 `xml:"Amt"`

	// Detailed balance information.
	DetailedBalance []*BalanceDetails6 `xml:"DtldBal,omitempty"`
}

func (b *BalanceDetails5) AddType() *BalanceType6Choice {
	b.Type = new(BalanceType6Choice)
	return b.Type
}

func (b *BalanceDetails5) SetUnrealised(value string) {
	b.Unrealised = (*Unrealised1Code)(&value)
}

func (b *BalanceDetails5) AddAmount() *AmountAndDirection31 {
	b.Amount = new(AmountAndDirection31)
	return b.Amount
}

func (b *BalanceDetails5) AddDetailedBalance() *BalanceDetails6 {
	newValue := new(BalanceDetails6)
	b.DetailedBalance = append(b.DetailedBalance, newValue)
	return newValue
}
