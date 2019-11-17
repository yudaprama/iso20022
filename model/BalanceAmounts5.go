package model

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type BalanceAmounts5 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection14 `xml:"HldgVal"`

	// Previous value of an individual financial instrument holding within a safekeeping account.
	PreviousHoldingValue *AmountAndDirection14 `xml:"PrvsHldgVal,omitempty"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection14 `xml:"BookVal,omitempty"`

	// Difference between holding value and the book value.
	UnrealisedGainLoss *AmountAndDirection14 `xml:"UrlsdGnLoss,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection14 `xml:"AcrdIntrstAmt,omitempty"`
}

func (b *BalanceAmounts5) AddHoldingValue() *AmountAndDirection14 {
	b.HoldingValue = new(AmountAndDirection14)
	return b.HoldingValue
}

func (b *BalanceAmounts5) AddPreviousHoldingValue() *AmountAndDirection14 {
	b.PreviousHoldingValue = new(AmountAndDirection14)
	return b.PreviousHoldingValue
}

func (b *BalanceAmounts5) AddBookValue() *AmountAndDirection14 {
	b.BookValue = new(AmountAndDirection14)
	return b.BookValue
}

func (b *BalanceAmounts5) AddUnrealisedGainLoss() *AmountAndDirection14 {
	b.UnrealisedGainLoss = new(AmountAndDirection14)
	return b.UnrealisedGainLoss
}

func (b *BalanceAmounts5) AddAccruedInterestAmount() *AmountAndDirection14 {
	b.AccruedInterestAmount = new(AmountAndDirection14)
	return b.AccruedInterestAmount
}
