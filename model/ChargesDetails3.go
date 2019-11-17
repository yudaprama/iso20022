package model

// Amount of money associated with a service.
type ChargesDetails3 struct {

	// Specifies the type of charges as a code or free text.
	Type *ChargesType1Choice `xml:"Tp"`

	// Specifies if it is a fixed amount or a percentage.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`
}

func (c *ChargesDetails3) AddType() *ChargesType1Choice {
	c.Type = new(ChargesType1Choice)
	return c.Type
}

func (c *ChargesDetails3) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	c.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return c.AmountOrPercentage
}
