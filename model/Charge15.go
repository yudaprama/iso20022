package model

// Amount of money associated with a service.
type Charge15 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType9Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`

	// Calculation basis for the charge or fee.
	CalculationBasis *CalculationBasis2Code `xml:"ClctnBsis,omitempty"`

	// Calculation basis for the charge or fee.
	ExtendedCalculationBasis *Extended350Code `xml:"XtndedClctnBsis,omitempty"`
}

func (c *Charge15) SetType(value string) {
	c.Type = (*ChargeType9Code)(&value)
}

func (c *Charge15) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge15) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge15) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge15) SetCalculationBasis(value string) {
	c.CalculationBasis = (*CalculationBasis2Code)(&value)
}

func (c *Charge15) SetExtendedCalculationBasis(value string) {
	c.ExtendedCalculationBasis = (*Extended350Code)(&value)
}
