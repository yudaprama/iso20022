package model

// Choice between balance, eligible balance and not eligible balance formats.
type BalanceFormat10Choice struct {

	// Provides information about balance related to a corporate action.
	Balance *SignedQuantityFormat8 `xml:"Bal"`

	// Provide eligible balance information in different formats.
	EligibleBalance *SignedQuantityFormat9 `xml:"ElgblBal"`

	// Provide not eligible balance information in different formats.
	NotEligibleBalance *SignedQuantityFormat9 `xml:"NotElgblBal"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat9 `xml:"FullPrdUnits"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat9 `xml:"PartWayPrdUnits"`
}

func (b *BalanceFormat10Choice) AddBalance() *SignedQuantityFormat8 {
	b.Balance = new(SignedQuantityFormat8)
	return b.Balance
}

func (b *BalanceFormat10Choice) AddEligibleBalance() *SignedQuantityFormat9 {
	b.EligibleBalance = new(SignedQuantityFormat9)
	return b.EligibleBalance
}

func (b *BalanceFormat10Choice) AddNotEligibleBalance() *SignedQuantityFormat9 {
	b.NotEligibleBalance = new(SignedQuantityFormat9)
	return b.NotEligibleBalance
}

func (b *BalanceFormat10Choice) AddFullPeriodUnits() *SignedQuantityFormat9 {
	b.FullPeriodUnits = new(SignedQuantityFormat9)
	return b.FullPeriodUnits
}

func (b *BalanceFormat10Choice) AddPartWayPeriodUnits() *SignedQuantityFormat9 {
	b.PartWayPeriodUnits = new(SignedQuantityFormat9)
	return b.PartWayPeriodUnits
}
