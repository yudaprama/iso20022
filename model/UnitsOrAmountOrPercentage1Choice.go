package model

// Choice between an amount or number of units or percentage.
type UnitsOrAmountOrPercentage1Choice struct {

	// Cash amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Number of units.
	Unit *DecimalNumber `xml:"Unit"`

	// Percentage of cash amount.
	Percentage *PercentageRate `xml:"Pctg"`
}

func (u *UnitsOrAmountOrPercentage1Choice) SetAmount(value, currency string) {
	u.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UnitsOrAmountOrPercentage1Choice) SetUnit(value string) {
	u.Unit = (*DecimalNumber)(&value)
}

func (u *UnitsOrAmountOrPercentage1Choice) SetPercentage(value string) {
	u.Percentage = (*PercentageRate)(&value)
}
