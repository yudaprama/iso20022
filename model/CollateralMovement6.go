package model

// Provides the collateral movement direction that is a delivery and optionaly a return.
type CollateralMovement6 struct {

	// Provides the collateral movement direction that is a delivery only.
	Deliver *Collateral8 `xml:"Dlvr"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral7 `xml:"Rtr,omitempty"`
}

func (c *CollateralMovement6) AddDeliver() *Collateral8 {
	c.Deliver = new(Collateral8)
	return c.Deliver
}

func (c *CollateralMovement6) AddReturn() *Collateral7 {
	c.Return = new(Collateral7)
	return c.Return
}
