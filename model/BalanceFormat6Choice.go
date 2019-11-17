package model

// Choice between balance, eligible balance and not eligible balance formats.
type BalanceFormat6Choice struct {

	// Provides information about balance related to a corporate action.
	Balance *SignedQuantityFormat7 `xml:"Bal"`

	// Provide eligible balance information in different formats.
	EligibleBalance *SignedQuantityFormat6 `xml:"ElgblBal"`

	// Provide not eligible balance information in different formats.
	NotEligibleBalance *SignedQuantityFormat6 `xml:"NotElgblBal"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat6 `xml:"FullPrdUnits"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat6 `xml:"PartWayPrdUnits"`
}

func (b *BalanceFormat6Choice) AddBalance() *SignedQuantityFormat7 {
	b.Balance = new(SignedQuantityFormat7)
	return b.Balance
}

func (b *BalanceFormat6Choice) AddEligibleBalance() *SignedQuantityFormat6 {
	b.EligibleBalance = new(SignedQuantityFormat6)
	return b.EligibleBalance
}

func (b *BalanceFormat6Choice) AddNotEligibleBalance() *SignedQuantityFormat6 {
	b.NotEligibleBalance = new(SignedQuantityFormat6)
	return b.NotEligibleBalance
}

func (b *BalanceFormat6Choice) AddFullPeriodUnits() *SignedQuantityFormat6 {
	b.FullPeriodUnits = new(SignedQuantityFormat6)
	return b.FullPeriodUnits
}

func (b *BalanceFormat6Choice) AddPartWayPeriodUnits() *SignedQuantityFormat6 {
	b.PartWayPeriodUnits = new(SignedQuantityFormat6)
	return b.PartWayPeriodUnits
}
