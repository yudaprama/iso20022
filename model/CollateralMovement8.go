package model

// Provides the collateral movement direction that is a delivery and optionaly a return.
type CollateralMovement8 struct {

	// Provides the collateral movement direction that is a delivery only.
	Deliver *Collateral12 `xml:"Dlvr"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral11 `xml:"Rtr,omitempty"`
}

func (c *CollateralMovement8) AddDeliver() *Collateral12 {
	c.Deliver = new(Collateral12)
	return c.Deliver
}

func (c *CollateralMovement8) AddReturn() *Collateral11 {
	c.Return = new(Collateral11)
	return c.Return
}
