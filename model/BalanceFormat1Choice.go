package model

// Choice between balance, eligible balance and not eligible balance formats.
type BalanceFormat1Choice struct {

	// Provides information about balance related to a corporate action.
	Balance *SignedQuantityFormat1 `xml:"Bal"`

	// Provide eligible balance information in different formats.
	EligibleBalance *SignedQuantityFormat2 `xml:"ElgblBal"`

	// Provide not eligible balance information in different formats.
	NotEligibleBalance *SignedQuantityFormat2 `xml:"NotElgblBal"`
}

func (b *BalanceFormat1Choice) AddBalance() *SignedQuantityFormat1 {
	b.Balance = new(SignedQuantityFormat1)
	return b.Balance
}

func (b *BalanceFormat1Choice) AddEligibleBalance() *SignedQuantityFormat2 {
	b.EligibleBalance = new(SignedQuantityFormat2)
	return b.EligibleBalance
}

func (b *BalanceFormat1Choice) AddNotEligibleBalance() *SignedQuantityFormat2 {
	b.NotEligibleBalance = new(SignedQuantityFormat2)
	return b.NotEligibleBalance
}
