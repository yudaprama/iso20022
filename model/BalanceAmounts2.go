package model

// Amounts linked to a securities balance, for example, holding value.
type BalanceAmounts2 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection6 `xml:"HldgVal"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection6 `xml:"BookVal,omitempty"`

	// Difference between holding value and the book value.
	UnrealisedGainLoss *AmountAndDirection6 `xml:"UrlsdGnLoss,omitempty"`
}

func (b *BalanceAmounts2) AddHoldingValue() *AmountAndDirection6 {
	b.HoldingValue = new(AmountAndDirection6)
	return b.HoldingValue
}

func (b *BalanceAmounts2) AddBookValue() *AmountAndDirection6 {
	b.BookValue = new(AmountAndDirection6)
	return b.BookValue
}

func (b *BalanceAmounts2) AddUnrealisedGainLoss() *AmountAndDirection6 {
	b.UnrealisedGainLoss = new(AmountAndDirection6)
	return b.UnrealisedGainLoss
}
