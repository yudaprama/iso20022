package model

// Amount of money associated with a service.
type ChargesDetails4 struct {

	// Specifies the type of charges as a code or free text.
	ChargesType *ChargesType1Choice `xml:"ChrgsTp"`

	// Amount of money asked or paid for the charge.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (c *ChargesDetails4) AddChargesType() *ChargesType1Choice {
	c.ChargesType = new(ChargesType1Choice)
	return c.ChargesType
}

func (c *ChargesDetails4) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}
