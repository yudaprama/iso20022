package model

// Amounts linked to a securities balance, for example, holding value.
type BalanceAmounts6 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection14 `xml:"HldgVal"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection14 `xml:"BookVal,omitempty"`

	// Difference between holding value and the book value.
	UnrealisedGainLoss *AmountAndDirection14 `xml:"UrlsdGnLoss,omitempty"`
}

func (b *BalanceAmounts6) AddHoldingValue() *AmountAndDirection14 {
	b.HoldingValue = new(AmountAndDirection14)
	return b.HoldingValue
}

func (b *BalanceAmounts6) AddBookValue() *AmountAndDirection14 {
	b.BookValue = new(AmountAndDirection14)
	return b.BookValue
}

func (b *BalanceAmounts6) AddUnrealisedGainLoss() *AmountAndDirection14 {
	b.UnrealisedGainLoss = new(AmountAndDirection14)
	return b.UnrealisedGainLoss
}
