package model

// Amounts linked to a securities balance, for example, holding value.
type BalanceAmounts3 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection6 `xml:"HldgVal,omitempty"`

	// Previous value of an individual financial instrument holding within a safekeeping account.
	PreviousHoldingValue *AmountAndDirection6 `xml:"PrvsHldgVal,omitempty"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection6 `xml:"BookVal,omitempty"`

	// Value of the position eligible for collateral purposes.
	EligibleCollateralValue *AmountAndDirection6 `xml:"ElgblCollVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection6 `xml:"AcrdIntrstAmt,omitempty"`
}

func (b *BalanceAmounts3) AddHoldingValue() *AmountAndDirection6 {
	b.HoldingValue = new(AmountAndDirection6)
	return b.HoldingValue
}

func (b *BalanceAmounts3) AddPreviousHoldingValue() *AmountAndDirection6 {
	b.PreviousHoldingValue = new(AmountAndDirection6)
	return b.PreviousHoldingValue
}

func (b *BalanceAmounts3) AddBookValue() *AmountAndDirection6 {
	b.BookValue = new(AmountAndDirection6)
	return b.BookValue
}

func (b *BalanceAmounts3) AddEligibleCollateralValue() *AmountAndDirection6 {
	b.EligibleCollateralValue = new(AmountAndDirection6)
	return b.EligibleCollateralValue
}

func (b *BalanceAmounts3) AddAccruedInterestAmount() *AmountAndDirection6 {
	b.AccruedInterestAmount = new(AmountAndDirection6)
	return b.AccruedInterestAmount
}
