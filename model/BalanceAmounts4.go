package model

// Amounts linked to a securities balance, for example, holding value.
type BalanceAmounts4 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection14 `xml:"HldgVal,omitempty"`

	// Previous value of an individual financial instrument holding within a safekeeping account.
	PreviousHoldingValue *AmountAndDirection14 `xml:"PrvsHldgVal,omitempty"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection14 `xml:"BookVal,omitempty"`

	// Value of the position eligible for collateral purposes.
	EligibleCollateralValue *AmountAndDirection14 `xml:"ElgblCollVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection14 `xml:"AcrdIntrstAmt,omitempty"`
}

func (b *BalanceAmounts4) AddHoldingValue() *AmountAndDirection14 {
	b.HoldingValue = new(AmountAndDirection14)
	return b.HoldingValue
}

func (b *BalanceAmounts4) AddPreviousHoldingValue() *AmountAndDirection14 {
	b.PreviousHoldingValue = new(AmountAndDirection14)
	return b.PreviousHoldingValue
}

func (b *BalanceAmounts4) AddBookValue() *AmountAndDirection14 {
	b.BookValue = new(AmountAndDirection14)
	return b.BookValue
}

func (b *BalanceAmounts4) AddEligibleCollateralValue() *AmountAndDirection14 {
	b.EligibleCollateralValue = new(AmountAndDirection14)
	return b.EligibleCollateralValue
}

func (b *BalanceAmounts4) AddAccruedInterestAmount() *AmountAndDirection14 {
	b.AccruedInterestAmount = new(AmountAndDirection14)
	return b.AccruedInterestAmount
}
