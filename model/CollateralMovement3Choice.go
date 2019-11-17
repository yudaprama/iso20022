package model

// Provides the collateral movement direction that is a delivery and optionaly a return, or a return only.
type CollateralMovement3Choice struct {

	// Provides the collateral movement direction that is a delivery and optionaly a return.
	CollateralMovementDirection *CollateralMovement6 `xml:"CollMvmntDrctn"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral7 `xml:"Rtr"`
}

func (c *CollateralMovement3Choice) AddCollateralMovementDirection() *CollateralMovement6 {
	c.CollateralMovementDirection = new(CollateralMovement6)
	return c.CollateralMovementDirection
}

func (c *CollateralMovement3Choice) AddReturn() *Collateral7 {
	c.Return = new(Collateral7)
	return c.Return
}
