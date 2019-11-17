package model

// Provides the agreed amount and the collateral movement direction.
type CollateralMovement5 struct {

	// Provides the call amount that is agreed and that will result in a collateral movement.
	AgreedAmount *ActiveCurrencyAndAmount `xml:"AgrdAmt"`

	// Provides the collateral movement direction that is a delivery and optionaly a return, or a return only.
	MovementDirection []*CollateralMovement3Choice `xml:"MvmntDrctn,omitempty"`
}

func (c *CollateralMovement5) SetAgreedAmount(value, currency string) {
	c.AgreedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralMovement5) AddMovementDirection() *CollateralMovement3Choice {
	newValue := new(CollateralMovement3Choice)
	c.MovementDirection = append(c.MovementDirection, newValue)
	return newValue
}
