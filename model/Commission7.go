package model

// Amount of money due to a party as compensation for a service.
type Commission7 struct {

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Service for which the commission is asked or paid.
	Type *CommissionType1 `xml:"Tp"`
}

func (c *Commission7) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission7) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission7) AddType() *CommissionType1 {
	c.Type = new(CommissionType1)
	return c.Type
}
