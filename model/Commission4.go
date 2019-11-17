package model

// Amount of money due to a party as compensation for a service.
type Commission4 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType1 `xml:"Tp"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *Commission4) AddType() *CommissionType1 {
	c.Type = new(CommissionType1)
	return c.Type
}

func (c *Commission4) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission4) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}
