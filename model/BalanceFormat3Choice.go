package model

// Choice between balance, eligible balance and not eligible balance formats.
type BalanceFormat3Choice struct {

	// Provides information about balance related to a corporate action.
	Balance *SignedQuantityFormat1 `xml:"Bal"`

	// Provide eligible balance information in different formats.
	EligibleBalance *SignedQuantityFormat2 `xml:"ElgblBal"`

	// Provide not eligible balance information in different formats.
	NotEligibleBalance *SignedQuantityFormat2 `xml:"NotElgblBal"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat2 `xml:"FullPrdUnits"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat2 `xml:"PartWayPrdUnits"`
}

func (b *BalanceFormat3Choice) AddBalance() *SignedQuantityFormat1 {
	b.Balance = new(SignedQuantityFormat1)
	return b.Balance
}

func (b *BalanceFormat3Choice) AddEligibleBalance() *SignedQuantityFormat2 {
	b.EligibleBalance = new(SignedQuantityFormat2)
	return b.EligibleBalance
}

func (b *BalanceFormat3Choice) AddNotEligibleBalance() *SignedQuantityFormat2 {
	b.NotEligibleBalance = new(SignedQuantityFormat2)
	return b.NotEligibleBalance
}

func (b *BalanceFormat3Choice) AddFullPeriodUnits() *SignedQuantityFormat2 {
	b.FullPeriodUnits = new(SignedQuantityFormat2)
	return b.FullPeriodUnits
}

func (b *BalanceFormat3Choice) AddPartWayPeriodUnits() *SignedQuantityFormat2 {
	b.PartWayPeriodUnits = new(SignedQuantityFormat2)
	return b.PartWayPeriodUnits
}
