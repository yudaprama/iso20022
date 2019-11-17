package model

// Provides the collateral movement direction that is a delivery and optionally a return.
type CollateralMovement11 struct {

	// Provides the collateral movement direction that is a delivery only.
	Deliver *Collateral16 `xml:"Dlvr"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral17 `xml:"Rtr,omitempty"`
}

func (c *CollateralMovement11) AddDeliver() *Collateral16 {
	c.Deliver = new(Collateral16)
	return c.Deliver
}

func (c *CollateralMovement11) AddReturn() *Collateral17 {
	c.Return = new(Collateral17)
	return c.Return
}
