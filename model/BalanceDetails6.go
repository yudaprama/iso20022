package model

// Balance related details for a portfolio.
type BalanceDetails6 struct {

	// Category of the financial asset balance type.
	Category *FinancialAssetTypeCategory1Code `xml:"Ctgy,omitempty"`

	// Balance type.
	Type *BalanceType7Choice `xml:"Tp,omitempty"`

	// Unrealised gain or loss.
	Unrealised *Unrealised1Code `xml:"Urlsd,omitempty"`

	// Balance amount.
	Amount *AmountAndDirection31 `xml:"Amt"`
}

func (b *BalanceDetails6) SetCategory(value string) {
	b.Category = (*FinancialAssetTypeCategory1Code)(&value)
}

func (b *BalanceDetails6) AddType() *BalanceType7Choice {
	b.Type = new(BalanceType7Choice)
	return b.Type
}

func (b *BalanceDetails6) SetUnrealised(value string) {
	b.Unrealised = (*Unrealised1Code)(&value)
}

func (b *BalanceDetails6) AddAmount() *AmountAndDirection31 {
	b.Amount = new(AmountAndDirection31)
	return b.Amount
}
