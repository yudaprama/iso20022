package model

// Amount of money due to a party as compensation for a service.
type Commission21 struct {

	// Type of commission.
	CommissionType *CommissionType5Choice `xml:"ComssnTp"`

	// Commission amount or commission rate applied.
	CommissionApplied *AmountOrRate3Choice `xml:"ComssnApld"`
}

func (c *Commission21) AddCommissionType() *CommissionType5Choice {
	c.CommissionType = new(CommissionType5Choice)
	return c.CommissionType
}

func (c *Commission21) AddCommissionApplied() *AmountOrRate3Choice {
	c.CommissionApplied = new(AmountOrRate3Choice)
	return c.CommissionApplied
}
