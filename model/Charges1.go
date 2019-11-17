package model

// Specifies the type of charge and the amont.
type Charges1 struct {

	// Type of charges.
	Type *ChargeType2FormatChoice `xml:"Tp,omitempty"`

	// Amount of charges.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (c *Charges1) AddType() *ChargeType2FormatChoice {
	c.Type = new(ChargeType2FormatChoice)
	return c.Type
}

func (c *Charges1) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}
