package model

// Choice between balance, eligible balance and not eligible balance formats.
type BalanceFormat7Choice struct {

	// Provides information about balance related to a corporate action.
	Balance *SignedQuantityFormat8 `xml:"Bal"`

	// Provide eligible balance information in different formats.
	EligibleBalance *SignedQuantityFormat9 `xml:"ElgblBal"`

	// Provide not eligible balance information in different formats.
	NotEligibleBalance *SignedQuantityFormat9 `xml:"NotElgblBal"`
}

func (b *BalanceFormat7Choice) AddBalance() *SignedQuantityFormat8 {
	b.Balance = new(SignedQuantityFormat8)
	return b.Balance
}

func (b *BalanceFormat7Choice) AddEligibleBalance() *SignedQuantityFormat9 {
	b.EligibleBalance = new(SignedQuantityFormat9)
	return b.EligibleBalance
}

func (b *BalanceFormat7Choice) AddNotEligibleBalance() *SignedQuantityFormat9 {
	b.NotEligibleBalance = new(SignedQuantityFormat9)
	return b.NotEligibleBalance
}
