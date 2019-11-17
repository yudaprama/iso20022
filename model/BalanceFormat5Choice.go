package model

// Choice between balance, eligible balance and not eligible balance formats.
type BalanceFormat5Choice struct {

	// Provides information about balance related to a corporate action.
	Balance *SignedQuantityFormat7 `xml:"Bal"`

	// Provide eligible balance information in different formats.
	EligibleBalance *SignedQuantityFormat6 `xml:"ElgblBal"`

	// Provide not eligible balance information in different formats.
	NotEligibleBalance *SignedQuantityFormat6 `xml:"NotElgblBal"`
}

func (b *BalanceFormat5Choice) AddBalance() *SignedQuantityFormat7 {
	b.Balance = new(SignedQuantityFormat7)
	return b.Balance
}

func (b *BalanceFormat5Choice) AddEligibleBalance() *SignedQuantityFormat6 {
	b.EligibleBalance = new(SignedQuantityFormat6)
	return b.EligibleBalance
}

func (b *BalanceFormat5Choice) AddNotEligibleBalance() *SignedQuantityFormat6 {
	b.NotEligibleBalance = new(SignedQuantityFormat6)
	return b.NotEligibleBalance
}
