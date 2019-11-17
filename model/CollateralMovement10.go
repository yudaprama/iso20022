package model

// Provides the agreed amount and the collateral movement direction.
type CollateralMovement10 struct {

	// Provides the call amount that is agreed and that will result in a collateral movement.
	AgreedAmount *ActiveCurrencyAndAmount `xml:"AgrdAmt"`

	// Provides the collateral movement direction that is a delivery and optionally a return, or a return only.
	MovementDirection []*CollateralMovement5Choice `xml:"MvmntDrctn,omitempty"`
}

func (c *CollateralMovement10) SetAgreedAmount(value, currency string) {
	c.AgreedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralMovement10) AddMovementDirection() *CollateralMovement5Choice {
	newValue := new(CollateralMovement5Choice)
	c.MovementDirection = append(c.MovementDirection, newValue)
	return newValue
}
