package model

// Total eligible balance for the corporate action and full and part way period units.
type TotalEligibleBalanceFormat1 struct {

	// Provides information about balance related to a corporate action.
	Balance *Quantity3Choice `xml:"Bal,omitempty"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat2 `xml:"FullPrdUnits,omitempty"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat2 `xml:"PartWayPrdUnits,omitempty"`
}

func (t *TotalEligibleBalanceFormat1) AddBalance() *Quantity3Choice {
	t.Balance = new(Quantity3Choice)
	return t.Balance
}

func (t *TotalEligibleBalanceFormat1) AddFullPeriodUnits() *SignedQuantityFormat2 {
	t.FullPeriodUnits = new(SignedQuantityFormat2)
	return t.FullPeriodUnits
}

func (t *TotalEligibleBalanceFormat1) AddPartWayPeriodUnits() *SignedQuantityFormat2 {
	t.PartWayPeriodUnits = new(SignedQuantityFormat2)
	return t.PartWayPeriodUnits
}
