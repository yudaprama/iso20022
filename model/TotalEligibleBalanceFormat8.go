package model

// Total eligible balance for the corporate action and full and part way period units.
type TotalEligibleBalanceFormat8 struct {

	// Provides information about balance related to a corporate action.
	Balance *Quantity17Choice `xml:"Bal,omitempty"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat6 `xml:"FullPrdUnits,omitempty"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat6 `xml:"PartWayPrdUnits,omitempty"`
}

func (t *TotalEligibleBalanceFormat8) AddBalance() *Quantity17Choice {
	t.Balance = new(Quantity17Choice)
	return t.Balance
}

func (t *TotalEligibleBalanceFormat8) AddFullPeriodUnits() *SignedQuantityFormat6 {
	t.FullPeriodUnits = new(SignedQuantityFormat6)
	return t.FullPeriodUnits
}

func (t *TotalEligibleBalanceFormat8) AddPartWayPeriodUnits() *SignedQuantityFormat6 {
	t.PartWayPeriodUnits = new(SignedQuantityFormat6)
	return t.PartWayPeriodUnits
}
