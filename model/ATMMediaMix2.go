package model

// Media mix selected.
type ATMMediaMix2 struct {

	// Number of notes or coins.
	Number *Number `xml:"Nb"`

	// Unit value.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal"`
}

func (a *ATMMediaMix2) SetNumber(value string) {
	a.Number = (*Number)(&value)
}

func (a *ATMMediaMix2) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}
