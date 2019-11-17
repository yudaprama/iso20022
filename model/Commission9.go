package model

// Amount of money due to a party as compensation for a service.
type Commission9 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType6Code `xml:"Tp"`

	// Service for which the commission is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Commission expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *Commission9) SetType(value string) {
	c.Type = (*CommissionType6Code)(&value)
}

func (c *Commission9) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Commission9) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission9) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}
