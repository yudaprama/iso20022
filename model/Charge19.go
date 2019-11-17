package model

// Amount of money associated with a service.
type Charge19 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType11Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *Charge19) SetType(value string) {
	c.Type = (*ChargeType11Code)(&value)
}

func (c *Charge19) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge19) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge19) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}
