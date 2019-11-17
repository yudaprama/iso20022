package model

// Amount of money associated with a service.
type Charge16 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType10Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *Charge16) SetType(value string) {
	c.Type = (*ChargeType10Code)(&value)
}

func (c *Charge16) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge16) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge16) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}
