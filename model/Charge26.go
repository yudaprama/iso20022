package model

// Choice of formats for the type of charge.
type Charge26 struct {

	// Type of charge.
	Type *ChargeType4Choice `xml:"Tp"`

	// Charge amount or charge rate applied.
	ChargeApplied *AmountOrRate3Choice `xml:"ChrgApld"`
}

func (c *Charge26) AddType() *ChargeType4Choice {
	c.Type = new(ChargeType4Choice)
	return c.Type
}

func (c *Charge26) AddChargeApplied() *AmountOrRate3Choice {
	c.ChargeApplied = new(AmountOrRate3Choice)
	return c.ChargeApplied
}
