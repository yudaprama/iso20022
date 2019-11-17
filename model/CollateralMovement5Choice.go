package model

// Provides the collateral movement direction that is a delivery and optionally a return, or a return only.
type CollateralMovement5Choice struct {

	// Provides the collateral movement direction that is a delivery and optionally a return.
	CollateralMovementDirection *CollateralMovement11 `xml:"CollMvmntDrctn"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral17 `xml:"Rtr"`
}

func (c *CollateralMovement5Choice) AddCollateralMovementDirection() *CollateralMovement11 {
	c.CollateralMovementDirection = new(CollateralMovement11)
	return c.CollateralMovementDirection
}

func (c *CollateralMovement5Choice) AddReturn() *Collateral17 {
	c.Return = new(Collateral17)
	return c.Return
}
