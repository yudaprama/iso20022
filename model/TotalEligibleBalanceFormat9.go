package model

// Total eligible balance for the corporate action and full and part way period units.
type TotalEligibleBalanceFormat9 struct {

	// Provides information about balance related to a corporate action.
	Balance *Quantity22Choice `xml:"Bal,omitempty"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat9 `xml:"FullPrdUnits,omitempty"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat9 `xml:"PartWayPrdUnits,omitempty"`
}

func (t *TotalEligibleBalanceFormat9) AddBalance() *Quantity22Choice {
	t.Balance = new(Quantity22Choice)
	return t.Balance
}

func (t *TotalEligibleBalanceFormat9) AddFullPeriodUnits() *SignedQuantityFormat9 {
	t.FullPeriodUnits = new(SignedQuantityFormat9)
	return t.FullPeriodUnits
}

func (t *TotalEligibleBalanceFormat9) AddPartWayPeriodUnits() *SignedQuantityFormat9 {
	t.PartWayPeriodUnits = new(SignedQuantityFormat9)
	return t.PartWayPeriodUnits
}
