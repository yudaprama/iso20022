package model

// Provides the collateral movement direction that is a delivery and optionaly a return, or a return only.
type CollateralMovement4Choice struct {

	// Provides the collateral movement direction that is a delivery and optionaly a return.
	CollateralMovementDirection *CollateralMovement8 `xml:"CollMvmntDrctn"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral11 `xml:"Rtr"`
}

func (c *CollateralMovement4Choice) AddCollateralMovementDirection() *CollateralMovement8 {
	c.CollateralMovementDirection = new(CollateralMovement8)
	return c.CollateralMovementDirection
}

func (c *CollateralMovement4Choice) AddReturn() *Collateral11 {
	c.Return = new(Collateral11)
	return c.Return
}
