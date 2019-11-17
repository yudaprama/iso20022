package model

// Amount of money associated with a service.
type Charge9 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType2 `xml:"Tp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Calculation basis for the charge or fee.
	CalculationBasis *CalculationBasis1 `xml:"ClctnBsis,omitempty"`
}

func (c *Charge9) AddType() *ChargeType2 {
	c.Type = new(ChargeType2)
	return c.Type
}

func (c *Charge9) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge9) AddCalculationBasis() *CalculationBasis1 {
	c.CalculationBasis = new(CalculationBasis1)
	return c.CalculationBasis
}
