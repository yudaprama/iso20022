package model

// Amount of money associated with a service.
type Charge11 struct {

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`

	// Type of service for which a charge is asked or paid.
	Type *ChargeType1 `xml:"Tp"`
}

func (c *Charge11) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge11) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge11) AddType() *ChargeType1 {
	c.Type = new(ChargeType1)
	return c.Type
}
