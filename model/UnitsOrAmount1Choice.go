package model

// Choice between an amount or number of units.
type UnitsOrAmount1Choice struct {

	// Currency and amount of the periodical payments.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Number of units to be subscribed or withdrawn.
	Unit *DecimalNumber `xml:"Unit"`
}

func (u *UnitsOrAmount1Choice) SetAmount(value, currency string) {
	u.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UnitsOrAmount1Choice) SetUnit(value string) {
	u.Unit = (*DecimalNumber)(&value)
}
