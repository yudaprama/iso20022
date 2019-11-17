package model

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type BalanceAmounts1 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection6 `xml:"HldgVal"`

	// Previous value of an individual financial instrument holding within a safekeeping account.
	PreviousHoldingValue *AmountAndDirection6 `xml:"PrvsHldgVal,omitempty"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection6 `xml:"BookVal,omitempty"`

	// Difference between holding value and the book value.
	UnrealisedGainLoss *AmountAndDirection6 `xml:"UrlsdGnLoss,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection6 `xml:"AcrdIntrstAmt,omitempty"`
}

func (b *BalanceAmounts1) AddHoldingValue() *AmountAndDirection6 {
	b.HoldingValue = new(AmountAndDirection6)
	return b.HoldingValue
}

func (b *BalanceAmounts1) AddPreviousHoldingValue() *AmountAndDirection6 {
	b.PreviousHoldingValue = new(AmountAndDirection6)
	return b.PreviousHoldingValue
}

func (b *BalanceAmounts1) AddBookValue() *AmountAndDirection6 {
	b.BookValue = new(AmountAndDirection6)
	return b.BookValue
}

func (b *BalanceAmounts1) AddUnrealisedGainLoss() *AmountAndDirection6 {
	b.UnrealisedGainLoss = new(AmountAndDirection6)
	return b.UnrealisedGainLoss
}

func (b *BalanceAmounts1) AddAccruedInterestAmount() *AmountAndDirection6 {
	b.AccruedInterestAmount = new(AmountAndDirection6)
	return b.AccruedInterestAmount
}
